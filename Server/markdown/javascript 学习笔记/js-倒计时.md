# js-倒计时

```
 function TimeBox(){
  let TimeBox = document.getElementsByTagName("time")[0];
  let toTime = new Date("4/14/2018 01:00:00");
  let onwTime = new Date();
  let countTime = toTime-onwTime;
  let hours = Math.floor(countTime/1000/60/60)
  let TmpMinutes = ((countTime/1000/60/60)-hours)
  let Minutes = Math.floor(TmpMinutes*60)
  let second = Math.floor(((TmpMinutes*60)-Minutes)*60)
  function add(name){
    if (name<10) {
        return "0"+name;
    }else{
        return name;
    }
  }
  TimeBox.innerHTML= add(hours)+"小时"+ add(Minutes)+"分"+ add(second)+"秒";
 }
 setInterval(TimeBox,1000);
```
