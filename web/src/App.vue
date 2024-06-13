<script setup>
import { RouterLink, RouterView } from 'vue-router'

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


</script>

<template>
<div class="container">
	<header>
		<div class="container">
			<div class="row align-items-center">
				<div class="col-lg-2 col-md-12 col-sm-12 col-xs-12">
					<a class="marble" href="/">
						<span class="use">ЗАЧЕМ</span> <span class="web">ВЫБИРАТЬ</span>
					</a>
				</div>
				<div class="col-lg-10 col-md-12 col-sm-12 col-xs-12">
					<ul class="choice justify-content-between d-flex">
						<li class="headerch"><a class="clickMenu" href="#procrast">Начало</a></li>
						<li class="headerch"><a class="clickMenu" href="#routes">Гарнир</a></li>
						<li class="headerch"><a class="clickMenu" href="#story">Основа</a></li>
						<li class="headerch"><a class="clickMenu" href="#final">Напиток</a></li>
						<li class="headerch"><a class="clickMenu" href="#order">Заказ</a></li>
					</ul>
				</div>
			</div>
		</div>
	</header>
    <RouterView>

    </RouterView>

    <footer>
    <router-link to="/"><img class="foot-img" src="./images/home.png" alt=""></router-link>
    <router-link to="/public"><img class="foot-img" src="./images/heart.png" alt=""></router-link>
    <router-link to="/order"><img class="foot-img" src="./images/bag-2.png" alt=""></router-link>
    <router-link to="/private"><img class="foot-img" src="./images/profile-circle.png" alt=""></router-link>
  </footer>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=ABeeZee&family=Heebo:wght@100..900&family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');
  
body {
	font-family: 'Open Sans', sans-serif;
	padding: 0;
	margin: 0;
	font-size: 13px;
	font-weight: 400;
	color: #2D3049;
	letter-spacing: 0.05em;
	background-image: url("../img/white-needle.png");
	overflow: hidden;
}
   /* body{ 
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .container{
    align-self: center;
    max-width: 1000px;
    margin-left: auto;
    margin-right: auto;
  }

  *{
    font-family: "Heebo", sans-serif;
  } */
  a {
	text-decoration: none;
  color: aliceblue;
}
  .clickMenu{
	width: 100px;
	height: 50px;
}
  .use {
    color: #aa064b;
}
.web {
    color: #429595; 
}
.headerch{
    list-style: none;
    float: left;
    text-align: center;
}
  .marble {
    display: block;
    font-size: 28px;
    color:  #b2aeb2;
    text-decoration: none;
    cursor: pointer;
}
  header {
    background-color: black;
	position: fixed;
	top: 0;
	left: 0;
	width: 100%;
}
  .footer-info{
    margin-top: 35px;
    margin-left: 100%;
    margin-right: 100%;
    color: rgb(33, 36, 61);
    width: 222px;
    font-family: Heebo;
    font-size: 14px;
    font-weight: 400;
    line-height: 21px;
    letter-spacing: 0px;
    text-align: center;
  }
  .images{
    display: flex;
    flex-direction: row;
    width: 230px;
    height: 30px;
    justify-content: space-between;
    margin-left: 100%;
    margin-right: 100%;
  }

  footer{
    position: fixed;
    bottom: 20%;

    display: flex;
    flex-direction: column;
    background-color: black;
    align-items: center;
    justify-content: center;


    gap:40px;
    padding: 20px;
    right: 20px;

    /* margin-left: 30%;
    margin-right: 40%; */

    border-radius: 40px;
  }
  .foot-img:hover{
    transform: scale(1.5);
    
  }

  /* header{
    padding: 15px;
    ul{
      list-style-type: none;
      display: flex;
      justify-content: end;
      gap: 25px;
    }
    a{
      text-decoration: none;
      color: black;

    } */

</style>
