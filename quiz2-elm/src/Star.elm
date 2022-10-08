module Star exposing (main)
import Html exposing (..)
import Html.Attributes exposing (class)
import Html.Events exposing (onClick)
import Browser
import Dict exposing (update)

type Msg = 
    Like | Unlike

initialModel : { liked : Bool }
initialModel =
    {
        
        liked = True
    }

update : Msg -> {  liked : Bool } -> { liked : Bool}
update msg model =
    case msg of
        Like ->
            { model | liked = False }
        Unlike ->
            { model | liked = True }
view : {  liked : Bool} -> Html Msg
view model =
    div []
    [ div [ class "header" ]
        [ h1[] [ text "PIKCHA" ] ]
    ,div [ class "content-flow" ]
        [ viewIcon model]
    
    ]
viewIcon: { liked: Bool } -> Html Msg
viewIcon  model =
    let 
        buttonType =
            if model.liked then "star"
            else
                "star_border"
        msg = if model.liked then Like
            else
                Unlike
    in
        div [ class "detailed-icon" ]
            [
            div [ class "like-button" ]
                    [ span 
                        [ class "material-icons md-48"
                        , onClick msg
                        ]
                        [ text buttonType ] ]
                
            ]



main : Program () {  liked : Bool} Msg
main =
    Browser.sandbox{
        init = initialModel
        ,view = view
        ,update = update
    }