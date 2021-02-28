# 利用伪元素让input更漂亮

```
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>利用伪元素修改单选框</title>
</head>
<style>
	* {
		margin: 0;
		padding: 0;
	}

	.container {
		width: 150px;
		margin: 0 auto;
		text-align: center;
		background: #ffffff;
	}

	.input-wrapper {
		position: relative;
		display: flex;
		flex-direction: row-reverse;
		align-items: center;
		justify-content: space-between;

	}

	input[type="radio"] {
		width: 20px;
		height: 20px;
		border-radius: 50%;
		outline-style:none;
		background: #ffffff;
		border: 0.5px solid silver;
		appearance: none;
		margin-bottom: 5px;
		-moz-appearance: none;
		-webkit-appearance: none;
	}

	input[type="radio"]+label::before {
		/* 找到input的下一个元素在之前插入伪元素*/
		content: '\a0';
		/* display: inline-block; */
		width: 20px;
		height: 20px;
		border-radius: 50%;
		background: #ffffff;
		border: 0.5px solid silver;
		/* 将第一行缩进 */
		/* text-indent: 2px; */
		line-height: 18px;
		/* 利用定位遮住原来的input */
		position: absolute;
		top: 0;
		right: 0;
	}
	input:checked+label::before {
		font: 500;
		content: '\2713';
		color: brown;
		background: ffffff;
	}
</style>

<body>
	<section class="container">
		<article>
			<form>
				<article class="input-wrapper ">
					<input type="radio" name="pay" checked="checked" id="WeChat" value="WeChat">
					<label for="WeChat">微信:</label>
				</article>
				<article class="input-wrapper">
					<input type="radio" name="pay" id="Alipay" value="Alipay">
					<label for="Alipay">支付宝:</label>
				</article>
			</form>
		</article>
	</section>
</body>

</html>
```
