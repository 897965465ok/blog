
import Home from "@/pages/home/home"
import { useEffect } from "react";
export default function Index(props) {

  const { title = "标题" } = props

  const handleClick = () => {
    console.log('点击事件');

  }

  useEffect(() => {
    console.log("useEffect")
  })

  return <>
    <h1>{title}</h1>
    <button onClick={handleClick}  >点击事件</button>
    <Home
      title={title}
    ></Home>
  </>
}