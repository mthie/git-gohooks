git-hooks
=========

Collection of git hooks. Most of the pre-commit hooks are Go specific.


usage
-----

* to install the hooks use:

		go get github.com/mthie/git-gohooks

		go get github.com/mthie/git-gohooks/githook-gobuild

		go get github.com/mthie/git-gohooks/githook-gofmt
		
		go get github.com/mthie/git-gohooks/githook-gotest

		
*	you have the choice to copy it to your .git/hooks/ folder or to symlink it with

		( cd .git/hooks && \

		  ln -s $GOPATH/bin/git-gohooks pre-commit && \

		  ln -s $GOPATH/bin/githook-gofmt pre-commit_01_gofmt && \

		  ln -s $GOPATH/bin/githook-gobuild pre-commit_02_gobuild && \ 
		  
		  ln -s $GOPATH/bin/githook-gotest pre-commit_03_gotest )

