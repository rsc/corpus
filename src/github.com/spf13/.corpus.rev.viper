commit 5ed0fc31f7f453625df314d8e66b9791e8d13003
Author: Kevin GEORGES <k4georges@gmail.com>
Date:   Tue Dec 13 10:38:49 2016 +0100

    Fix MergeInConfig error return
    
    UnsupportedConfigError was returned if config file not found
    
    * Swap getConfigFile and getConfigType call
    * Add a unit test
