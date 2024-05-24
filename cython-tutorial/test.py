import pvh

file_path = "/Users/nicth99p/bin/projects/PredVirusHost/sequences/protein/phage_fasta/fasta-headers.txt"

d = pvh.load_protein_names(file_path)
d = pvh.protein_count(d)
print(d)

