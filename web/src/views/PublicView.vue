<script>
export default {
  data() {
    return {
      users: []
    };
  },
  mounted() {
    this.fetchUsers();
  },
  methods: {
    fetchUsers() {
      fetch('https://dummyjson.com/users')
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(data => {
          this.users = data.users;
        })
        .catch(error => {
          console.error('Fetch error:', error);
        });
    }
  }
};
</script>

<template>
    <div class="users-layout">
      <div>
        <div v-for="user in users" :key="user.id" class="single-user">
          <p class="name">{{ user.firstName + ' ' + user.lastName  + ' ' + user.maidenName}}</p>
          <p class="mail">{{ user.email }}</p>
        </div>
    </div>
    </div>
  </template>

<style scoped>
    .users-layout{
        padding-top: 70px;
    }
    .single-user{
        border-radius: 10px;
        box-shadow: 0px 2.75px 9px 0px rgba(0, 0, 0, 0.19),0px 0.25px 3px 0px rgba(0, 0, 0, 0.04);

        display: flex;
        flex-direction: row;
        justify-content: space-around;

        width: 402px;
        height: 56px;

        margin-left: auto;
        margin-right: auto;
        margin-bottom: 18px;
    }

    .name{
        font-weight: 600;
        align-self: center;
    }
    .mail{
        color: rgb(133, 133, 133);
        font-size: small;
        align-self: center;
    }
</style>
  
