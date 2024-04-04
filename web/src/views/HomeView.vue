<script>
    import { RouterLink, RouterView } from 'vue-router'
    export default{

    }
</script>

<template>
<div class="home-layout">
        <div class="text-field">
            <h1>Home page</h1>
            

            <RouterLink to ='/login'>
                <button>
                Authorization
                </button>
            </RouterLink>
        </div>
        <img src="../images/home.png"/>   
</div>

</template>

<style scoped>
    img{
        height: 300px;
    }

    button{
        border-radius: 2px;
        background: rgb(255, 100, 100);
        color: rgb(255, 255, 255);
        font-family: Heebo;
        font-size: 20px;
        font-weight: 500;
        line-height: 29px;
        letter-spacing: 0px;
        text-align: center;

        /* Rectangle 10 */

        width: 208px;
        height: 43px;
        outline: none;
        border: none;

    }

    .home-layout{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 500px;
    }

    .text-field{
        display: flex;
        flex-direction: column;
        gap: 30px;
    }
    h1{
        color: rgb(33, 36, 61);
        font-family: Heebo;
        font-size: 44px;
        font-weight: 700;
        line-height: 60px;
        letter-spacing: 0px;
        text-align: left;
        margin: 0;
    }
    p{
        color: rgb(33, 36, 61);
        font-family: Heebo;
        font-size: 16px;
        font-weight: 400;
        line-height: 24px;
        letter-spacing: 0px;
        text-align: left;
        width: 500px;
    }
</style>