commit ded73eae5db7e7a0ef6f55aace87a2873c5d2b74
Author: Ugorji Nwoke <ugorji@gmail.com>
Date:   Sat Jan 7 08:32:03 2017 -0500

    codec: only generate Selfer methods if AST doesn't have those methods already
    
    This is a cleaner solution, that checking using reflection if
    codec.Selfer is implemented.
    
    It solves the issue where a type implements codec.Selfer implicitly (by
    embedding).  In this situation, we still want to generate Selfer
    implementations for the enclosing type.
    
    We can't do this with reflection effectively, because golang reflection
    has the limitation that you cannot (at runtime) determine whether a Type
    explicitly/directly (itself) implements an interface or whether it
    implements that interface implicitly (via embedding).
    
    By design (of codecgen)
    - first (initial) phase of codecgen is the only place where we
      inspect the AST to infer the types to codecgen on.
    - an execution of codecgen only works on files within a
      specific directory (within a single package).
      It makes the problem easier to handle and solve.
    
    However, this implementation has some caveats:
    - only generate types if the type doesn't explicitly implement Selfer methods
      in one of the files passed to that execution of codecgen
    
    we can check this while doing the AST analysis for type names to generate
    in this mode, we don't check anymore at the top of Gen(...) to see if the reflect.Type implements codec.Selfer
    
    This will handle the following cases:
    
    - file1.go: defines type T, and CodecEncodeSelfer and CodecDecodeSelfer
      run codecgen on file1.go
    - file1.go: defines type T
      file2.go: defines CodecEncodeSelfer and CodecDecodeSelfer for type T
      run codecgen on file1.go AND file2.go (in one execution)
    
    However, it will NOT handle the following cases:
    - file1.go: defines type T
      file2.go: defines CodecEncodeSelfer and CodecDecodeSelfer for type T
      run codecgen on file1.go (without passing in file2.go also)
    
    In practice, folks can easily work around this limitation.
    
    Fixes #185 #186
