#!/usr/bin/env python3

"""Generate enums and instruction structs based on the official registry.

Requires `spirv.core.grammar.json` in the same folder.
You can fetch the source grammar at the following URL:
https://www.khronos.org/registry/spir-v/
"""

import json
from pathlib import Path
from collections import defaultdict
import re

folder = Path(__file__).parent
infile = folder / "spirv.core.grammar.json"
outdir = folder


def new_file(name):
    outdir.mkdir(parents=True, exist_ok=True)
    fp = Path(outdir, name).open('w')
    fp.write("// This file was generated. DO NOT EDIT!\n\n")
    fp.write("package spirv\n\n")
    return fp


def main():
    with open(infile) as fp:
        data = json.load(fp)

    # TODO metainfo/constants (version)
    write_operands(data)
    write_instr_opcodes(data)


def write_enum(fp, enum):
    type_ = enum['kind']
    fp.write(f"type {type_} uint32\n")
    fp.write(f"func (v {type_}) Verify() error {{ return nil }}\n")

    fp.write("const (\n")
    for enumerant in enum['enumerants']:
        name = field_name(enumerant['enumerant'])
        fp.write(f"    {type_}{name} {type_} = {enumerant['value']}\n")
    fp.write(")\n\n")


LITERAL_MAP = {
    'LiteralString': 'String',
    # TOCHECK most of these are actually context-dependant
    'LiteralExtInstInteger': 'uint32',
    'LiteralSpecConstantOpInteger': 'uint32',
    'LiteralInteger': 'uint32',
    'LiteralContextDependentNumber': 'uint32',
}


def get_type(type_):
    if type_ in LITERAL_MAP:
        return LITERAL_MAP[type_]
    elif type_.startswith('Id'):
        return 'Id'
    else:
        return type_


def write_operands(data):
    # Split by category
    types = defaultdict(list)
    for operand in data['operand_kinds']:
        types[operand['category']].append(operand)

    with new_file("enums_gen.go") as fp:
        for operand in types.pop('BitEnum'):
            write_enum(fp, operand)
        for operand in types.pop('ValueEnum'):
            write_enum(fp, operand)

    with new_file("types_gen.go") as fp:
        # We only use one externally defined Id type
        types.pop('Id')

        for operand in types.pop('Literal'):
            kind = operand['kind']
            if kind not in LITERAL_MAP:
                print("Unknown kind:", kind)

        for operand in types.pop('Composite'):
            # Need to assign a name for structs with the same types
            fields = "; ".join(f"{chr(0x41 + i)} {get_type(ftype)}"
                               for i, ftype in enumerate(operand['bases']))
            fp.write(f"type {operand['kind']} struct {{ {fields} }}\n\n")

    if types:
        print("Unvisited types:", sorted(types))


def field_name(orig_name):
    # "name" : "'Luma Intra Partition Mask'"
    # "name" : "'Offset'"
    # "name" : "'D~ref~'"
    # "name" : "'Vector 1'"
    # "enumerant" : "1D"
    # "name" : "'Argument 0', +\n'Argument 1', +\n..."
    # "name" : "'Variable, Parent, ...'"
    if "\n" in orig_name or orig_name == "'Variable, Parent, ...'":
        return "Argv"

    with_ws = orig_name.strip("'")
    no_ws = re.sub(r" +(.)", lambda m: m.group(1).upper(), with_ws)
    alphanum = re.sub(r"\W", "_", no_ws)

    # special cases
    if alphanum == "Opcode":
        alphanum = "OpcodeParam"
    elif alphanum == "IdResult":
        alphanum = "ResultId"
    elif alphanum == "IdResultType":
        alphanum = "ResultType"
    return alphanum


def write_instr_opcodes(data):
    # These opcodes accept multiple literals but don't state so in the grammer
    EXTRA_ARGV = {'OpDecorate', 'OpMemberDecorate', 'OpExecutionMode', 'OpLoopMerge'}
    # 'OpConstant', 'OpSpecConstant' # these are actually a LiteralContextDependentNumber type

    def const_name(opname):
        return f"opcode{opname[2:]}"

    def instructions_filter(inst_data):
        """Skip some reserved or weird instructions."""
        return (
            inst_data['class'] not in ('Reserved', '@exclude')
            and not inst_data['opname'].endswith("GOOGLE")
        )

    instructions = list(filter(instructions_filter, data['instructions']))

    with new_file("opcodes_gen.go") as fp:
        fp.write("const (\n")
        for inst_data in instructions:
            fp.write(f"    {const_name(inst_data['opname'])} = {inst_data['opcode']}\n")
        fp.write(")\n")

    with new_file("instructions_gen.go") as fp:
        for inst_data in instructions:
            fp.write(f"type {inst_data['opname']} struct {{\n")

            # Fields
            fields = set()
            for operand in inst_data.get('operands', []):
                name = field_name(operand.get('name', operand['kind']))
                if name in fields:
                    name += "2"  # naive name collision prevention
                fields.add(name)
                type_ = operand['kind']
                type_ = get_type(type_)
                if 'quantifier' in operand:
                    if operand['quantifier'] == '*':
                        type_ = f'[]{type_}'
                    type_ = f'{type_} `spirv:"optional"`'
                fp.write(f"    {name} {type_}\n")
            if inst_data['opname'] in EXTRA_ARGV:
                fp.write("    Argv []uint32\n")
            fp.write("}\n")

            # Methods
            fp.write(f"func (op *{inst_data['opname']}) Opcode() uint32"
                     f" {{ return {const_name(inst_data['opname'])} }}\n")
            optional = str(inst_data['class'] == 'Debug').lower()
            fp.write(f"func (op *{inst_data['opname']}) Optional() bool {{ return {optional} }}\n")
            fp.write(f"func (op *{inst_data['opname']}) Verify() error {{ return nil }}\n")
            fp.write("\n")

        fp.write("func init() {\n")

        for inst_data in instructions:
            # The old package had this, but idk why?
            # optionals = {}
            # fields = set()
            # for operand in inst_data.get('operands', []):
            #     name = field_name(operand['name']) if 'name' in operand else operand['kind']
            #     if name in fields:
            #         name += "2"  # naive name collision prevention
            #     if operand.get('quantifier') == '?':
            #         # All current defaults are integers
            #         optionals[name] = "0"

            # optional_str = ", ".join(f"{k}: {v}" for k, v in optionals.items())
            # fp.write(f"    bind(func() Instruction {{ return &{inst_data['opname']}"
            #          f"{{{optional_str}}} }})\n")
            fp.write(f"    bind(func() Instruction {{ return &{inst_data['opname']}{{}} }})\n")
        fp.write("}\n")


if __name__ == '__main__':
    main()
