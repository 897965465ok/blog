# setTimeout应用移入弹出子菜单

```
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Document</title>
	<style>
		* {
			padding: 0;
			margin: 0;
			list-style: none;
		}

		header,
		nav {
			height: 50px;
			width: 1200px;
			margin: 0 auto;
			background: rgb(63, 63, 63);
			color: white;
		}


		nav>ul>li {
			position: relative;
			height: 100%;
			line-height: 100%;
			flex: 1 0 20%;
			display: flex;
			justify-content: center;
			align-items: center;
			cursor: pointer
		}

		nav ul {
			width: 100%;
			height: 100%;
			display: flex;
			align-items: center;
			justify-content: space-around;
		}

		.sub-menu {
			display: none;
			position: absolute;
			width: 100%;
			background: rgb(63, 63, 63);
		}

		.sub-menu>ul {
			width: 100%;
			height: 100%;
			display: flex;
			flex-direction: column;
			justify-content: space-around;
		}

		.sub-menu>ul>li,
		.sub-menu>ul>li>a {
			color: white;
			text-decoration: none;
			width: 100%;
			height: 40px;
			line-height: 40px;
			text-align: center;
		}

		nav>ul>li:hover,
		.sub-menu>ul>li:hover {
			background: rgb(47, 47, 47);
		}

		.container {
			width: 500px;
			height: 200px;
			border: rgb(47, 47, 47) solid 1px;
			margin: 0 auto;
			margin-top: 10%;
			position: relative;
		}
	</style>
</head>

<body>
	<header>
		<nav>
			<ul>
				<li>
					战争
				</li>
				<li>
					恐怖
				</li>
				<li>
					僵尸

				</li>
				<li>
					休闲

				</li>
				<li>
					沙盒
				</li>
			</ul>
		</nav>
	</header>
	<section class="container">
		<div class="sub-menu">
			<ul>
				<li><a href="#">战地6</a></li>
				<li><a href="#">使命召唤16</a></li>
				<li><a href="#">真三国无双8</a></li>
				<li><a href="#">三国:全面战争</a></li>
				<li><a href="#">文明6</a></li>
			</ul>
		</div>
		<div class="sub-menu">
			<ul>
				<li><a href="#">恶灵附身2</a></li>
				<li><a href="#">逃生2</a></li>
				<li><a href="#">毁灭战士4</a></li>
				<li><a href="#">黎明杀机</a></li>
				<li><a href="#">死亡空间3</a></li>
			</ul>
		</div>
		<div class="sub-menu">
			<ul>
				<li><a href="#">生化危机6</a></li>
				<li><a href="#">往日不在</a></li>
				<li><a href="#">僵尸世界大战</a></li>
				<li><a href="#">亿万僵尸</a></li>
				<li><a href="#">最后的生还者2</a></li>
			</ul>
		</div>
		<div class="sub-menu">
			<ul>
				<li><a href="#">辐射:避难所</a></li>
				<li><a href="#">模拟人生</a></li>
				<li><a href="#">杀戮尖塔</a></li>
				<li><a href="#">王国之心</a></li>
				<li><a href="#">超级马里奥</a></li>
			</ul>
		</div>
		<div class="sub-menu">
			<ul>
				<li><a href="#">塞尔达传说</a></li>
				<li><a href="#">巫师3</a></li>
				<li><a href="#">合金装备5</a></li>
				<li><a href="#">看门狗</a></li>
				<li><a href="#">狂怒2</a></li>
			</ul>
		</div>
	</section>
</body>
<script>
	let mainMenus = document.querySelectorAll('header>nav>ul>li');
	let subMenus = document.querySelectorAll('.sub-menu');
	let timer = [];

	mainMenus.forEach((item, index) => {
		item.onmouseover = () => {
			subMenus[index].style.display = 'block';
		}
		item.onmouseout = () => {
			timer[index] = setTimeout(() => {
				subMenus[index].style.display = 'none';
			}, 500)
		}
	})

	subMenus.forEach((item, index) => {
		item.onmouseover = () => {
			clearTimeout(timer[index]);
		}
		item.onmouseout = () => {
			timer[index] = setTimeout(() => {
				subMenus[index].style.display = 'none';
			}, 500)
		}
	})


</script>

</html>
```
