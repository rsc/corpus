# Go Corpus v0.01

This Git repository is a first draft of a common corpus of Go code to be used
for evaluating possible changes to Go, Go libraries, and Go tools.

The actual code is not in this repo. Instead, after checkout you must
run `gitrestore` to do all the necessary git checkouts.
That script uses metadata stored in .corpus.* files throughout
the tree. The effect is like git submodules, but I couldn't get git submodules
to work well enough to use them directly.
I wanted to store the actual code in the repo, but it got a bit too big.

The script `addproject` attempts to add a new project to the repo,
meaning the code required for x/... (and its dependencies) for some x.
See the git history for the list of projects in the repo.
They can also be identified by files matching .corpus.project.*.

This is version 0.01 of the corpus. Each of the projects has all the code
needed to resolve imports, but there has been no check that the code
actually builds.

**If you are only interested in analyzing the set of Go source code
defined by the corpus, don't use the Git repository. Instead,
download [https://storage.googleapis.com/go-corpus/go-corpus-0.01.tar.gz](https://storage.googleapis.com/go-corpus/go-corpus-0.01.tar.gz).**

