module TestSignumMain where
import Prelude
import TestSignum (x)
import Effect.Console (log)
import Effect (Effect)

main :: Effect Unit
main = do
  log $ show (1.0 / x)
