<script>
import { RouterLink, RouterView } from 'vue-router'
export default{

}

</script>

<template>
<div class="container">
    <header>
      <ul>
        <RouterLink to ='/'>
          Home
        </RouterLink>
        <RouterLink to ='/public'>
          Users
        </RouterLink>
        <RouterLink to ='/private'>
          Profile
        </RouterLink>
      </ul>
    </header>
    <RouterView>

    </RouterView>

    <footer>
    <router-link to="/"><img class="foot-img" src="./images/home.png" alt=""></router-link>
    <router-link to="/public"><img class="foot-img" src="./images/heart.png" alt=""></router-link>
    <router-link to="/private"><img class="foot-img" src="./images/bag-2.png" alt=""></router-link>
    <router-link to="/private"><img class="foot-img" src="./images/profile-circle.png" alt=""></router-link>
  </footer>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=ABeeZee&family=Heebo:wght@100..900&family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');
  
  body{
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
    bottom: 10%;

    display: flex;
    flex-direction: row;
    background-color: black;
    align-items: center;
    justify-content: center;

    height: 80px;
    width: 350px;

    gap:40px;

    margin-left: 30%;
    margin-right: 40%;

    border-radius: 40px;
  }
  .foot-img:hover{
    transform: scale(1.5);
    
  }

  header{
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

    }
  }
</style>
