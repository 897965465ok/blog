import { useContext, useEffect } from 'react';
import { StoreContext } from 'redux-react-hook';
import { useHistory } from "react-router-dom"
import http from "../../http"
export default function () {
    // const store = useContext(StoreContext);
    // let history = useHistory()
    const getUserInfo = async () => {
        let articles = await http.getArticles()
        console.log(articles)
    }
    useEffect(async () => {
        // await getUserInfo()
    })
    return (
        <div>
           
        </div>
    )
}

