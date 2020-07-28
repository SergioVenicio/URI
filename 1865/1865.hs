main = do
    input <- getLine
    let times = read input :: Int
    getHero(times)

splitWhen :: (Char -> Bool) -> String -> [String]
splitWhen char sep = 
    case dropWhile char sep of
        "" -> []
        s' -> word : splitWhen char s''
            where (word, s'') = break char s'

toString :: Char -> String
toString c = [c]

getHero repeat = do
    hero <- getLine
    let heroName = splitWhen (==' ') hero
    if head heroName == "Thor"
        then
        
            putStrLn "Y" :: IO()
        else 
            putStrLn "N" :: IO()

    if repeat - 1 > 0
        then
            getHero(repeat-1)
        else
            return ""
