import sys

if len(sys.argv) < 4:
   print "Usage: "+sys.argv[0]+" <input> <output> <port>\r\n"
   sys.exit()

fd = open(sys.argv[1], "r")
nfd = open(sys.argv[2], "w")

file_content = fd.readlines()
file_content = [x.rstrip() for x in file_content]
fd.close()

for line in file_content:
    line.strip("\r").strip("\n")
    out_buffer = line + ":" + str(sys.argv[3])
    nfd.write(out_buffer+"\r\n")

nfd.close()
print "Finished!"
sys.exit()
