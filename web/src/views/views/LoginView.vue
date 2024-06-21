<script setup>
import { ref } from 'vue'
import axios from 'axios';

const email = ref('')
const password = ref('')
const error = ref('')

const login = async () => {
  try {
    const response = await axios.post('/login', {
      email: email.value,
      password: password.value
    }, {
      headers: { 'Content-Type': 'application/json' }
    });

    if (response.data.success) {
      const data = await response.data;
      error.value = data.message || 'Успешный вход'
      setTimeout(() => {
        window.location.href = '/private'; 
      }, 3000);
    } else {
      const data = await response.data;
      error.value = data.message || 'Неверный логин или пароль';
    }

  } catch (err) {
    error.value = 'Произошла ошибка при отправке запроса';
    console.error(err);
  }
}
</script>

<template>
    <div class="login-layout">
      <form @submit.prevent="login">
        <div class="form-head">
            <h3 class="title">Authorization</h3><p v-if="error" class="title" style="color: red;">{{ error }}</p>
        </div>
          <label for="email">Email</label>
          <input type="text" id="email" v-model="email" required>
          <label for="password">Password</label>
          <input type="password" id="password" v-model="password" required>
        <button type="submit">Submit</button>
        <router-link to="/registrate">go registrate</router-link>
      </form>
    </div>
  </template>

  <style scoped>
  .login-layout{
    display: flex;
    flex-direction: column;
    justify-content: center;
    height: 500px;
    padding-top: 10vh
  }
  .title{height: 24px; margin-top: 16px;}
  .form-head{
    display: flex;
    flex-direction: row;
    justify-content: space-between;

    margin-bottom: 33px;
    height: 24px;
  }
  form{
    display: flex;
    flex-direction: column;

    /* Inside Auto Layout */
    flex: none;
    order: undefined;
    flex-grow: 0;
    margin-left: auto;
    margin-right: auto;

    width: 309px;
    height: 277px;

    background: rgb(255, 255, 255);

    border-radius: 10px;
    box-shadow: 0px 2.75px 9px 0px rgba(0, 0, 0, 0.19),0px 0.25px 3px 0px rgba(0, 0, 0, 0.04);

    padding-right: 13px;
    padding-left: 13px;
  }
    label{
        color: rgba(0, 0, 0, 0.6);
        font-size: 14px;
        font-weight: 400;
        line-height: 20px;
        letter-spacing: 0.25px;
        text-align: left;
        align-self: center;
        width: 152px;
        margin-top: 13px;
    }
    input{
        border-radius: 10px;
        border: none;
        box-shadow: 0px 2.75px 9px 0px rgba(0, 0, 0, 0.19),0px 0.25px 3px 0px rgba(0, 0, 0, 0.04);
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-self: center;
        padding: 0px;
        width: 126px;
        height: 29px;

        padding: 0 10px 0 10px;
    }
    button{
        color: rgba(0, 0, 0, 0.6);
        font-size: 14px;
        font-weight: 400;
        line-height: 20px;
        letter-spacing: 0.25px;
        text-align: center;
        border: none;
        width: 60px;

        align-self: center;
        background: rgb(255, 255, 255);

        margin-top: 20px;
        margin-bottom: 43px;
    }
    a{
      width: 100%;
      text-align: center;
      color: blue;
      margin-bottom: 25px;
    }
  </style>
  
