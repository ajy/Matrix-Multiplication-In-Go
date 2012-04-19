import random
import argparse

d = []
def argparser():
    parser = argparse.ArgumentParser(description='Parser')
    parser.add_argument("--size",type=int,nargs=1)
    global d
    d = parser.parse_args()
 
def num_gen():
    for i in range(d.size[0]):
        for j in range(d.size[0]-1):
            print(int(random.random()*10000),end=',',sep = '')
        print(int(random.random()*10000))

if __name__=="__main__":
    argparser()
    num_gen()
