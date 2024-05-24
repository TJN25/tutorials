import logging as log

def load_protein_names(f):
    with open(f, 'r') as f:
        lines = f.readlines()
    proteins_dict = {}
    for line in lines:
        try:
            value, key = line.strip().split(' ')
            value = value[1:]
        except ValueError:
            continue
        if not key in proteins_dict:
            proteins_dict[key] = [value]
        else:
            proteins_dict[key].append(value)
    return proteins_dict

def protein_count(d):
    for key, value in d.items():
        length = len(value)
        d[key] = [value, length]
    log.info(d)

verbose = True
if verbose:
    log.basicConfig(format="%(message)s", level=log.DEBUG)
    log.info("Verbose flag used: printing output.")
else:
    log.basicConfig(format="%(levelname)s: %(message)s")
