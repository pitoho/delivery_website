import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import PrivateView from '../views/PrivateView.vue'
import PublicView from '../views/PublicView.vue'
import FrontView from '../views/FrontView.vue'
import FullMenu from '@/views/FullMenu.vue'
import OrderView from '@/views/OrderView.vue'

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
      component: FrontView
    },
    {
      path: '/private',
      name: 'private',
      component: PrivateView
    },
    {
      path: '/public',
      name: 'public',
      component: FullMenu
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/front',
      name: 'front',
      component: FrontView
    },
    {
      path: '/order', 
      name: 'order',
      component: OrderView
    }
  ],
  scrollBehavior() {
    // Всегда прокручивает к началу страницы
    return { top: 0 };
  }
})

router.beforeEach((to, from, next)=>{
  if ((to.path==='/private' && !checkAuth()) || (to.path==='/order' && !checkAuth())){
    next('/login')
  }else{
    next()
  }
})

export default router
