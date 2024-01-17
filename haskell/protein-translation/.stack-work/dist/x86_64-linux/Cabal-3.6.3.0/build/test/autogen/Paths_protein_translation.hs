{-# LANGUAGE CPP #-}
{-# LANGUAGE NoRebindableSyntax #-}
{-# OPTIONS_GHC -fno-warn-missing-import-lists #-}
{-# OPTIONS_GHC -w #-}
module Paths_protein_translation (
    version,
    getBinDir, getLibDir, getDynLibDir, getDataDir, getLibexecDir,
    getDataFileName, getSysconfDir
  ) where


import qualified Control.Exception as Exception
import qualified Data.List as List
import Data.Version (Version(..))
import System.Environment (getEnv)
import Prelude


#if defined(VERSION_base)

#if MIN_VERSION_base(4,0,0)
catchIO :: IO a -> (Exception.IOException -> IO a) -> IO a
#else
catchIO :: IO a -> (Exception.Exception -> IO a) -> IO a
#endif

#else
catchIO :: IO a -> (Exception.IOException -> IO a) -> IO a
#endif
catchIO = Exception.catch

version :: Version
version = Version [1,1,1,3] []

getDataFileName :: FilePath -> IO FilePath
getDataFileName name = do
  dir <- getDataDir
  return (dir `joinFileName` name)

getBinDir, getLibDir, getDynLibDir, getDataDir, getLibexecDir, getSysconfDir :: IO FilePath



bindir, libdir, dynlibdir, datadir, libexecdir, sysconfdir :: FilePath
bindir     = "/mnt/data/dev/local/exercism/haskell/protein-translation/.stack-work/install/x86_64-linux/16e1b02ae1ef92495043565a5b82f8c1d69f078768ec481e9704d4f73f215c9e/9.2.7/bin"
libdir     = "/mnt/data/dev/local/exercism/haskell/protein-translation/.stack-work/install/x86_64-linux/16e1b02ae1ef92495043565a5b82f8c1d69f078768ec481e9704d4f73f215c9e/9.2.7/lib/x86_64-linux-ghc-9.2.7/protein-translation-1.1.1.3-UWJEIkrnXcIvPrEnoQYm9-test"
dynlibdir  = "/mnt/data/dev/local/exercism/haskell/protein-translation/.stack-work/install/x86_64-linux/16e1b02ae1ef92495043565a5b82f8c1d69f078768ec481e9704d4f73f215c9e/9.2.7/lib/x86_64-linux-ghc-9.2.7"
datadir    = "/mnt/data/dev/local/exercism/haskell/protein-translation/.stack-work/install/x86_64-linux/16e1b02ae1ef92495043565a5b82f8c1d69f078768ec481e9704d4f73f215c9e/9.2.7/share/x86_64-linux-ghc-9.2.7/protein-translation-1.1.1.3"
libexecdir = "/mnt/data/dev/local/exercism/haskell/protein-translation/.stack-work/install/x86_64-linux/16e1b02ae1ef92495043565a5b82f8c1d69f078768ec481e9704d4f73f215c9e/9.2.7/libexec/x86_64-linux-ghc-9.2.7/protein-translation-1.1.1.3"
sysconfdir = "/mnt/data/dev/local/exercism/haskell/protein-translation/.stack-work/install/x86_64-linux/16e1b02ae1ef92495043565a5b82f8c1d69f078768ec481e9704d4f73f215c9e/9.2.7/etc"

getBinDir     = catchIO (getEnv "protein_translation_bindir")     (\_ -> return bindir)
getLibDir     = catchIO (getEnv "protein_translation_libdir")     (\_ -> return libdir)
getDynLibDir  = catchIO (getEnv "protein_translation_dynlibdir")  (\_ -> return dynlibdir)
getDataDir    = catchIO (getEnv "protein_translation_datadir")    (\_ -> return datadir)
getLibexecDir = catchIO (getEnv "protein_translation_libexecdir") (\_ -> return libexecdir)
getSysconfDir = catchIO (getEnv "protein_translation_sysconfdir") (\_ -> return sysconfdir)




joinFileName :: String -> String -> FilePath
joinFileName ""  fname = fname
joinFileName "." fname = fname
joinFileName dir ""    = dir
joinFileName dir fname
  | isPathSeparator (List.last dir) = dir ++ fname
  | otherwise                       = dir ++ pathSeparator : fname

pathSeparator :: Char
pathSeparator = '/'

isPathSeparator :: Char -> Bool
isPathSeparator c = c == '/'
