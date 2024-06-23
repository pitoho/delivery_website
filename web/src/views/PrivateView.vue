<script setup>
import { useBucketStore } from '@/stores/BucketStore';
import { onMounted, ref } from 'vue';

const bucketStore = useBucketStore()
let user = ref([])
let orders = ref([])

onMounted(
    user.value = bucketStore.getCookie("user").slice(1,-1).split(","),
    // console.log(user.value),
    orders.value = JSON.parse(decodeURIComponent(bucketStore.getCookie("orders")).replaceAll('+',' ')),
    // console.log(orders.value)
)
</script>

<template>
    <div>
      <div class="user-layout">
        <h1 class="error" v-if="error">{{ error }}</h1>
        <div class="user-content" v-if="user">
            <h2>My profile</h2>
            <div class="info">
                <div>
                    <p>Username:{{ user[0] }}</p>
                    <p>Lastnane:{{ user[1] }}</p>
                    <p>Phone:{{ user[2] }}</p>
                    <p>Email:{{ user[3] }}</p>
                </div>
                    <img :src="'https://www.kindpng.com/picc/m/150-1503949_computer-icons-user-profile-male-profile-icon-png.png'" alt="pfp">
            </div>

        </div>
      </div>
    </div>
  </template>

<style scoped>
    .user-layout{
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        height: 500px;
        margin-top: 15vh
    }
    .error{
        color: rgb(33, 36, 61);
        font-family: Heebo;
        font-size: 44px;
        font-weight: 700;
        line-height: 60px;
        letter-spacing: 0px;
        
        margin-left: auto;
        margin-right: auto;

        width: fit-content;

        align-self: center;

        margin-top: 70px;
    }
    .user-content{
        height: 500px;
        width: 600px;

        margin-left: auto;
        margin-right: auto;
    }
    .info{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }

    img{
        height: 243px;
        width: 243px;
        border-radius: 50%;
    }

    h2{
        color: rgb(33, 36, 61);
        font-family: Heebo;
        font-size: 44px;
        font-weight: 700;
        line-height: 60px;
        letter-spacing: 0px;
        text-align: left;
    }

    p{
        border-radius: 10px;
        box-shadow: 0px 2.75px 9px 0px rgba(0, 0, 0, 0.19),0px 0.25px 3px 0px rgba(0, 0, 0, 0.04);
        width:fit-content;
        text-align: left;

        padding: 10px 30px 10px 10px;

        font-weight: 600;
    }
</style>
  
