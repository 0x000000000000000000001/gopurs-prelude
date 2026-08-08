module TestNegZero where
import Prelude
import Effect.Console (log)
main = log (show (1.0 / (-0.0)))
