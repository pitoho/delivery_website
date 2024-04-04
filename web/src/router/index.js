import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import PrivateView from '../views/PrivateView.vue'
import PublicView from '../views/PublicView.vue'

function checkAuth(){
  if (localStorage.getItem('token')){
    return true
  } else{
    return false
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/private',
      name: 'private',
      component: PrivateView
    },
    {
      path: '/public',
      name: 'public',
      component: PublicView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    }
  ]
})

router.beforeEach((to, from, next)=>{
  if (to.path==='/private' && !checkAuth()){
    next('/login')
  }else{
    next()
  }
})

export default router
