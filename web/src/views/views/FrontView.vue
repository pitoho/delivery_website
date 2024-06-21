
<script setup>
import ProductItem from '@/components/ProductItem.vue';
import { useBucketStore } from '@/stores/BucketStore';
const bucketStore = useBucketStore()

const anchors = document.querySelectorAll('a[href*="#"]')

for (let anchor of anchors) {
  anchor.addEventListener('click', function (e) {
    e.preventDefault()

    const blockID = anchor.getAttribute('href').substr(1)

    document.getElementById(blockID).scrollIntoView({
      behavior: 'smooth',
      block: 'start'
    })
  })
}

function addToOrder(tag){
	bucketStore.addRandomItem(tag)
}

</script>

<template>

	<div id="block-content"></div>


	<section id="travel" class="travel">
		<div class="container">
			<div class="offer">
				<h1 class="offer__title">
					Положитесь на удачу
				</h1>
				<p class="offer__intro">
					Вам лень выбирать?<br />
					Всего 3 шага, и заказ уже в пути!:
				</p>
				<a href="#footer" class="offer__btn btn">
					Узнать больше
				</a>
			</div>
			<a href="#procrast" class="arrow-down"></a>
		</div>
	</section>
	<section id="procrast" class="procrast">
		<div class="container">
			<div class="offer">
				<h2>
					Шаг 1
				</h2>
				<h1 class="offer__title">
					Поверьте нам!
				</h1>
				<p class="offer__intro">
					Вам практически ничего не нужно
				</p>
			</div>
			<a href="#routes" class="arrow-down"></a>
		</div>
	</section>
	<section id="routes" class="routes">
		<div class="container">
			<h2>
				Шаг 2
			</h2>
			<h1 class="routes__title">
				Гарнир
			</h1>
			<h3>Чем бы закусить?</h3>
			<div class="destinations">
				<div class="destinations__block">
					<div class="destinations__img">
						<img src="https://images.unsplash.com/photo-1606755456206-b25206cde27e?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
							alt="Пофигист">
					</div>
					<div class="destinations__wrapper">
						<h3 class="destinations__title">
							Жареное
						</h3>
						<p class="destinations__desc">
							Картофель фри, гренки и тд
						</p>
						<div class="destinations__info-wrap">
							<p class="destinations__price">
								О да!
							</p>
							<a href="#story" @click="addToOrder(3)" class="destinations__btn">

							</a>
						</div>
					</div>
				</div>
				<div class="destinations__block">
					<div class="destinations__img">
						<img src="https://images.unsplash.com/photo-1600490036275-35f5f1656861?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" alt="Индия">
					</div>
					<div class="destinations__wrapper">
						<h3 class="destinations__title">
							Вареное
						</h3>
						<p class="destinations__desc">
							Пюре или макарошки
						</p>
						<div class="destinations__info-wrap">
							<p class="destinations__price">
								Мммм...
							</p>
							<a href="#story" @click="addToOrder(3)" class="destinations__btn">

							</a>
						</div>
					</div>
				</div>
				<div class="destinations__block">
					<div class="destinations__img">
						<img src="https://images.unsplash.com/photo-1525604499593-d92408d2adbc?q=80&w=2473&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" alt="Индия">
					</div>
					<div class="destinations__wrapper">
						<h3 class="destinations__title">
							Не буду
						</h3>
						<p class="destinations__desc">
							Кому какое дело?
						</p>
						<div class="destinations__info-wrap">
							<p class="destinations__price">
								Дальше
							</p>
							<a href="#story" class="destinations__btn">
							</a>
						</div>
					</div>
				</div>
			</div>
		</div>
	</section>
	<section id="story" class="story">
			<div class="container">
				<h2>
					Шаг 3
				</h2>
				<h1 class="story__title">
					Основное блюдо
				</h1>
				<p class="offer__intro">
					Чем порадуете себя сегодня?
				</p>
		</div>
		<div class="story-content" id ="story-content">
			<div class="container container_full-height">
				<ul class="story-content__list">
					<li class="story-content__item item1">
						<b class ="bold">Рыба/Море</b><br /> От лосося до мидии
						<a href="#final" class="arrow-down arr1" @click="addToOrder(1)"></a>
					</li>
					<li class="story-content__item item2">
						<b class ="bold">Мясо/Птица</b><br />Хрустящая курочка или стейк?
						<a href="#final" class="arrow-down arr2" @click="addToOrder(1)"></a>
					</li>
					<li class="story-content__item item3">
						<b class ="bold">Веган</b><br /> Салатикииии
						<a href="#final" class="arrow-down arr3" @click="addToOrder(1)"></a>
					</li>
				</ul>
				<div class="PictureStory" id="PictureStory">
				</div>
			</div>
		</div>
	</section>
	<section id="final" class="final">
		<div class="container finals">
			<h2>
				Шаг 4
			</h2>
			<h1 class="final__title">
				Напиток
			</h1>

	<div class="final-content" id ="final-content">
		<p class = "text__final">Пить или не пить?<br />
		Вот в чем вопрос!</p>
		<div class = "wrapper">
			<!-- <div class = "card_menu"></div> -->
			<div class = "card_menu card1"><a href = "#order" class ="wraptext">Не буду</a></div>
			<div class = "card_menu card2"><a href = "#order" class ="wraptext" @click="addToOrder(2)">Чай</a></div>
			<div class = "card_menu big card3"><a href = "#order" class ="wraptext" @click="addToOrder(2)">Алкоголь</a></div>
			<div class = "card_menu med card4"><a href = "#order" class ="wraptext" @click="addToOrder(2)">Газировка</a></div>
		</div>
	</div>
</div>
</section>

<section id="order" class="story">
			<div class="container">
				<h2>
					Заказ
				</h2>
				<h1 class="story__title">
					Основываясь на ваших предпочтениях:
				</h1>
				<p class="offer__intro">
					
				</p>
		</div>
		<div class="story-content" id ="story-content">
			<div class="container container_full-height">
				<ProductItem v-for="product in bucketStore.order" :key="product.id_dish" :product="product" />
			</div>
		</div>
		<RouterLink to="/order" class="offer__btn btn">
			В корзину
		</RouterLink>
</section>

</template>

<style scoped>

/* -----------------------------------------------------------------------Шапочка--------------------------------------------------------------------  */


* {
    margin: 0;
    padding: 0;
}



.logo p {
    margin-left: 20px;
    margin-top: 2px;
    margin-bottom: 5px;
    font-family:  "Exo2-Medium", sans-serif;
    font-size: 13px;
    color: white;
}

.choice li a {
    display: block;
    width: max-content;
    font-size: 17px;
    color: white;
	padding: 20px;
    cursor: pointer;
}
.choice li a:hover {
    transition: all 0.5s ease;
    -webkit-transition: all 0.5s ease;
    -o-transition: all 0.5s ease;
    -ms-transition:all 0.5s ease;
    -moz-transition:all 0.5s ease;
    background: #00ADE4;
}
.block-top-auth {
    width: 100px;
    float: left;
    margin-left: 40px;
    height: 80px;
}
.block-top-auth p {
    margin-top: 12px;
    text-align: center;
    margin-bottom: 8px;
    
}
.block-top-auth a {
    text-decoration: none;
    color: white;
    font-family:  "Exo2-Medium", sans-serif;
    font-size: 15px;
    text-decoration: underline;
    cursor: pointer;
}
.block-top-auth a:hover {
    transition: all 0.5s ease;
    -webkit-transition: all 0.5s ease;
    -o-transition: all 0.5s ease;
    -ms-transition:all 0.5s ease;
    -moz-transition:all 0.5s ease;
    color: #00ADE4;
}
section {
	padding-top: 144px;
	height: 100vh;
}
* {
	box-sizing: border-box;
}
ul, li {
	display: block;
	padding: 0;
	margin: 0;
}

h1, h2, h3, h4, h5, h6 {
	padding: 0;
	margin: 0;
	text-transform: uppercase; /*чтобы заголовки были большими*/
	font-weight: 800;
	
}
h1{
	background-color: rgba(0, 0, 0, 0.13);
	
}
p {
	margin: 0;
}

a {
	text-decoration: none;
}

.container {
	width: 100%;
	margin: 0 auto;
}

.container_full-height {
	height: 100%;
	display: flex;
	align-items: center;
}

.btn {
	padding: 18px 0;
	background: #FF3773;
	color: #fff;
	font-size: 16px;
	font-weight: 800;
	display: block;
	width: 244px;
	margin: 0 auto;
	transition: all 0.5s ease; /*- чтобы при наведении на кнопку, цвет кнопки менялся плавно*/
	text-align: center;
}
.btn:hover {
	background: #FF185E;
	transform: scale(1.5,1.5);
}

.btn-ghost {
	padding: 18px 0;
	color: #FF3773;
	border: 1px solid #FF3773;
	font-size: 16px;
	font-weight: 800;
	display: block;
	width: 329px;
	text-align: center;
	transition: all 0.5s ease; /*- чтобы при наведении на кнопку, цвет кнопки менялся плавно*/
}

.btn-ghost:hover {
	color: #fff;
	background: #FF3773;
	border: 1px solid #FF3773;
}

.section-title {
	font-size: 30px;
	margin-bottom: 89px;
}
/* Экран смерти */
.death {
	color: white;
	background-color: black;
    height: 100vh;
}
.reaper img{
	object-fit: fill;
	width: 30%;
	margin-left: 36%;
}

/*Первая-вторая секция*/

.travel {
	background: linear-gradient(224.92deg, rgba(19, 19, 51, 0.18) 17.19%, rgba(40, 39, 74, 0.81) 81.74%), url('https://images.unsplash.com/photo-1504674900247-0877df9cc836?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D') no-repeat center top / cover; /*cover - нужно чтобы при больших разрешениях чем наша картинка, растягивалась на всю ширину разрешения*/
	text-align: center;
	color: #fff;
	padding: 77px 0 80px;
	height: 100vh;
}

.offer {
	margin-top: 195px;
}

.offer__title {
	font-size: 60px;
	letter-spacing: 0.22em;
	line-height: 1.35em;
}

.offer__intro {
	font-size: 20px;
	margin-top: 10px;
	margin-bottom: 34px;
}

.offer__btn {
	margin: 0 auto;
}

.arrow-down {
	background: url(../img/arrow-down.svg) no-repeat bottom;
    width: 50px;
    height: 50px;
    display: block;
    margin: 239px auto 0;
    opacity: 0.5;
    transition: all 0.5s ease;

}

.arrow-down:hover {
	opacity: 1;
}


.procrast{
	text-align: center;
	color: #fff;
	padding: 77px 0 80px;
	animation-name: picchange;
    animation-duration: 8s;
    animation-iteration-count: infinite;
	height: 100vh;
}
@keyframes picchange{
0%{
	background: url("https://images.unsplash.com/photo-1536521642388-441263f88a61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80") no-repeat center top / cover;
}
12.5%{
	background: url("https://images.unsplash.com/photo-1536521642388-441263f88a61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80") no-repeat center top / cover;
}
37.5%{
	background: url('https://images.unsplash.com/photo-1445311429009-0b7e0cff714a?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1374&q=80') no-repeat center top / cover; 
}
62.5%{
	background: url('https://images.unsplash.com/photo-1493711662062-fa541adb3fc8?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80') no-repeat center top / cover; 
}
87.5%{
	background: url('https://images.unsplash.com/photo-1453227588063-bb302b62f50b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80') no-repeat center top / cover;
}
100%{
	background: url("https://images.unsplash.com/photo-1536521642388-441263f88a61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80") no-repeat center top / cover;
}
}

/*Третья секция*/

.destinations {
	
	display: flex;
	justify-content: space-between;
	margin-bottom: 89px;
	margin-top: 20px;
}

.destinations__block {
	width: 331px;
	border: 1px solid #DCDCDC;
	transition: all 0.5s ease;
margin-left: 10px;
height: auto;
}

.destinations__img {

}

.destinations__img img {
	object-fit: fill;
	width: 100%;
	height: 358px;
}

.destinations__wrapper {
	padding: 52px 35px 47px;
}

.destinations__title {
	font-size: 18px;
	margin-bottom: 16px;
}

.destinations__desc {
	line-height: 1.67em;
}

.destinations__info-wrap {
	margin-top: 38px;
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.destinations__price {
	color: #FF3773;
	font-size: 20px;
}

.destinations__btn {
	background: url('../img/arrow-next.svg');
	width: 25px;
	height: 21px;
	display: block;
	opacity: 0;
	transition: all 0.5s ease;
}

.destinations__block:hover {
	box-shadow: 0px 15px 50px rgba(0, 0, 0, 0.13);

}

.destinations__block:hover .destinations__btn {
	opacity: 1;
}
.routes__title{
	font-size: 60px;
	letter-spacing: 0.22em;
	line-height: 1.35em;
	background-color: rgba(196, 210, 247, 0);
}

/*Четвертая секция*/

.story-content {
	padding: 20px;
	background: #F1F1F5 no-repeat right / 55%;
	height: 600px;
}

.story-content__item {
	font-size: 16px;
	width: 401px;
	position: relative;
	padding-left: 22px;
	line-height: 1.01em;
	margin-bottom: 74px;
	line-height:20px
}

.story-content__item:before {
	content:'';
	display: block;
	width: 7px;
	height: 40px;
	background: #9E9EB7;
	position: absolute;
	left: 0;
}
.PictureStory{
	background: url('https://images.unsplash.com/photo-1614548539924-5c1f205b3747?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D') no-repeat center bottom / cover ;
	border-radius:10px;
	height: 50vh;
	width: 100vw;
}
.PictureStory:hover{
	animation-name: storch;
    animation-duration: 60s;
    animation-iteration-count: 1;
	animation-fill-mode: forwards;
}
.arr1,.arr2,.arr3{
	background: url('../img/arrow-down.svg') no-repeat center bottom/ contain;
	border-radius: 3px;
	background-color: rgb(0,0,0,0.13);
	margin: 5px auto 0;
height: 40px;
width: 40px;
}
/* .arr1:hover+.PictureStory{
	animation-name: storch;
    animation-duration: 2s;
    animation-iteration-count: infinite; -------------- Не получилось, объекты не соседние(
}
@keyframes storch{
	100%{
	background: url('../img/india.jpg')
	}
} */

ul.story-content__list:hover+.PictureStory {
	animation-name: storch;
    animation-duration: 60s;
    animation-iteration-count: 1;
	animation-fill-mode: forwards;
}
.item1::before{
	background: rgb(0, 132, 51);
}
.arr1:hover{
	background-color: rgb(0, 132, 51);
}
.item2::before{
	background: rgb(0, 26, 111);
}
.arr2:hover{
	background-color: rgb(0, 26, 111);
}
.item3::before{
	background: rgb(114, 0, 0);
}
.arr3:hover{
	background-color: rgb(114, 0, 0);
}
.story__title {
	font-size: 60px;
	letter-spacing: 0.22em;
	line-height: 1.35em;
	background-color: rgba(196, 210, 247, 0);
}
@keyframes storch{
	2%{
		background: url('https://images.unsplash.com/photo-1525755662778-989d0524087e?q=80&w=2574&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D') no-repeat center bottom / cover 
	}
	100%{
		background: url('https://images.unsplash.com/photo-1614548539924-5c1f205b3747?q=80&w=2370&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D') no-repeat center bottom / cover 
	}
}
.bold{
	font-size: 20px;
}
/* -------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
/* Пятая секция (финал) */
.container .finals{
	display: flex;
	flex-direction: column;
}
.final__title{
	font-size: 60px;
	letter-spacing: 0.15em;
	line-height: 1.35em;
	background-color: rgba(196, 210, 247, 0);
	margin-bottom: 20px;
}
.text__final{
	text-align: center;
	margin: 0 auto;
	margin-bottom: 30px;
	font-size: 25px;
}
.final-content{
	background: #F1F1F5 no-repeat right / 55%;
	height: 60vh;
}
.card_menu{
    position: relative;
    background-color:forestgreen;
    width: 90%;
    padding: 10px;
    border-radius: 40px;
    box-sizing: border-box;
    margin: 10px;

    font-family: Verdana, Geneva, Tahoma, sans-serif;
    align-content: center;
    text-align: center;
	color: transparent;
	transition: all 0.5s ease; /*- чтобы при наведении на кнопку, цвет кнопки менялся плавно*/
	display: flex;

}
.wrapper{
	background-color: aquamarine;
    height: 41vh;
    display: grid;
    grid-template-columns: repeat(4,auto);
    grid-template-rows: repeat(2,auto);
    background-color: transparent;
}
.card1{
	background: url("https://images.unsplash.com/photo-1525604499593-d92408d2adbc?q=80&w=2473&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D") no-repeat center / cover;
}
.card1:hover{
	color: black;
	font-size: 30px;
	transform: scale(1.1,1.1);
    background-blend-mode: screen;
	background: linear-gradient(to right, #91919f,#8d8d9a), url("https://images.unsplash.com/photo-1525604499593-d92408d2adbc?q=80&w=2473&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D") no-repeat center / cover;
}
.card1:hover .wraptext{
	opacity: 1;
}
.card2{
background: url(https://images.unsplash.com/photo-1594631252845-29fc4cc8cde9?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D3) no-repeat bottom / cover;
}
.card2:hover{
	color: black;
	font-size: 30px;
	transform: scale(1.1,1.1);
	background: linear-gradient(to right, #91919f,#8d8d9a), url(https://images.unsplash.com/photo-1594631252845-29fc4cc8cde9?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D) no-repeat bottom / cover;
    background-blend-mode: screen;
    background-size: cover;
}
.card2:hover .wraptext{
	opacity: 1;
}
.card4{
	background: url(https://images.unsplash.com/photo-1556710807-a9261921475f?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D) no-repeat center / cover;
}
.card4:hover{
	color: black;
	font-size: 30px;
	transform: scale(1.1,1.1);
	background: linear-gradient(to right, #91919f,#8d8d9a), url(https://images.unsplash.com/photo-1556710807-a9261921475f?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D) no-repeat center / cover;
    background-blend-mode: screen;
    background-size: cover;
}
.card4:hover .wraptext{
	opacity: 1;
}
.card3{
	background: url(https://images.unsplash.com/photo-1547595628-c61a29f496f0?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D) no-repeat center / cover;
}
.card3:hover{
	color: black;
	font-size: 30px;
	transform: scale(1.1,1.1);
	background: linear-gradient(to right, #91919f,#8d8d9a), url(https://images.unsplash.com/photo-1547595628-c61a29f496f0?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D) no-repeat center / cover;
    background-blend-mode: screen;
    background-size: cover;
}
.card3:hover .wraptext{
	opacity: 1;
}
.med{
	grid-column: span 2;
	width: 94%;
}
.big{
	grid-column: span 2;
	grid-row: span 2;
	width: 100%;
	height: 42vh;
}
.wraptext{
	color: black;
	opacity: 0;
	margin: 0 auto;
	align-self: center;
	font-style:unset;
}
/* -------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
/* Футер */
/* .footer{
	height: 870px;
}
.footer_content{
    background: #F1F1F5 no-repeat right / 55%;
    display: flex;
    margin-top: 144px;
    height: 100%;
	border-radius: 30px;
	align-items: center;
	justify-content: space-between;
}
.triang{
    width: 50%;
    height: 500px;
	background: 
		linear-gradient(to right, #F1F1F5, transparent), 
		url(../img/YADEDINSAID.jpg) no-repeat  bottom / cover ;
	background-blend-mode: screen;
	background-size: cover;
	border-radius: 20px;
	right: 0;
}

.description{
	font-size: 16px;
	width: 401px;
	position: relative;
	padding-left: 22px;
	line-height: 1.01em;
	margin-bottom: 74px;
	line-height:20px;
	margin: 20px;
}
.arrPrav{
	margin: 40px;
}
.footer_item{
	font-size: 16px;
	width: 401px;
	position: relative;
	padding-left: 22px;
	line-height: 1.01em;
	margin-bottom: 50px;
	line-height:20px
} */
.fi1::before{
	content:'';
	display: block;
	width: 40px;
	height: 40px;
	background: url(https://www.svgrepo.com/show/96440/up-arrow.svg) no-repeat center /cover;
	transform: rotate(180deg);
	position: absolute;
	left: -30px;
	border-radius: 10px;
}
.fi1:hover::before{
    animation: downpoint 1.5s ease-out infinite;
}
@keyframes downpoint{
    from {
        transform: translateY(-10px) rotate(-180deg);
    }

    50% {
        transform: translateY(10px) rotate(-180deg);
    }

    to {
        transform: translateY(-10px) rotate(-180deg);
    }
}
.fi2::before{
	content:'';
	display: block;
	width: 40px;
	height: 40px;
	background: url(https://svgsilh.com/svg/296395.svg) no-repeat center /cover;
	position: absolute;
	left: -30px;
	border-radius: 10px;
}
.fi2:hover::before{
	animation: shake 0.3s ease-out infinite;
}
@keyframes shake{
    from {
        transform: translateY(-3px) rotate(-30deg);
    }

    50% {
        transform: translateY(3px) rotate(30deg);
    }

    to {
        transform: translateY(-3px) rotate(-30deg);
    }
}
.fi3::before{
	content:'';
	display: block;
	width: 40px;
	height: 40px;
	background: url(../img/Tilda_Icons_5_Event_Agency/Tilda_Icons_5ev_magic.svg) no-repeat center /cover;
	position: absolute;
	left: -30px;
	border-radius: 10px;
}
.fi3:hover::before{
	animation: switch 1s cubic-bezier(1,-0.07,1,.52) infinite ;
}
@keyframes switch{
100%{
	transform: rotate(360deg);
}
}
/* -------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
/* Медиа для регулировки размера */
@media screen and (max-width: 1200px) {
	
}
@media screen and (max-width: 992px) {
	.marble {
		text-align: center;
	}
	.offer__title {
		font-size: 40px;
	}
	.wrapper{
		display: flex;
		flex-wrap: wrap;
	}
	.story{
		height: 130vh;
	}
	.final{
		height: 130vh;
	}
	.final-content{
		height: 800px;
	}
	.container .finals{
		height: 130vh;
	}
	.footer_content{
		display: flex;
		flex-wrap: wrap;
	}
}

@media screen and (max-width: 768px) {
.choice{
	display:flex;
	flex-direction:row;
	flex-wrap:wrap;
	margin: 0 auto;
}
.clickMenu{
	width: 100px;
	height: 50px;
}
.destinations{
	display:flex;
	flex-direction:row;
	flex-wrap:wrap;
}
.container_full-height{
	display:flex;
	flex-direction:row;
	flex-wrap:wrap;
}
.story-content__list{
	margin: 0 auto;
}
.story__title{
	font-size: 40px;
}
.final__title{
	font-size: 40px;
}
.offer__intro{
	font-size: 20px;
}
.bold{
	font-size: 20px;
}
.story-content__item{
	font-size: 20px;
}
}
@media screen and (max-width: 576px) {
.offer__title{
font-size: 35px;
}
.routes__title{
font-size: 35px;
}
.story__title{
	font-size: 20px;
}
.final__title{
	font-size: 20px;
}
.offer__intro{
	font-size: 15px;
}
.bold{
	font-size: 13px;
}
.story-content__item{
	font-size: 12px;
}
body{
	overflow:scroll;
}
}
@media screen and (max-width: 385px) {
.offer__title{
font-size: 15px;
}
.btn{
	margin-left:0px;
}
.arrPrav{
	margin-left:-15px
}
.footer_item{
	font-size: 10px;
}
.opis{
	font-size: 10px;
	margin-left: -35px;
}
.description{
	width: 280px;
}
}

</style>