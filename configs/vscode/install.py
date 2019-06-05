src = '''// Installing github.com/mdempsky/gocode FAILED
// Installing github.com/uudashr/gopkgs/cmd/gopkgs SUCCEEDED
// Installing github.com/ramya-rao-a/go-outline FAILED
// Installing github.com/acroca/go-symbols FAILED
// Installing golang.org/x/tools/cmd/guru FAILED
// Installing golang.org/x/tools/cmd/gorename FAILED
// Installing github.com/go-delve/delve/cmd/dlv SUCCEEDED
// Installing github.com/stamblerre/gocode FAILED
// Installing github.com/rogpeppe/godef SUCCEEDED
// Installing golang.org/x/tools/cmd/goimports FAILED
// Installing golang.org/x/lint/golint FAILED
// Installing golang.org/x/tools/cmd/gopls FAILED'''

import os
def install():
    pkgs = get_pkgs()
    res = list(map(lambda x: os.popen('go install '+ x).read(), pkgs))
    print(list(res))
    print(pkgs)

def get_pkgs():
    return list(map(lambda x : x.split('Installing ')[1].split(' ')[0], src.split('\n')))

if __name__ == '__main__':
    install()