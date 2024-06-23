import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import RegistrateView from '../views/RegistrateView.vue'
import PrivateView from '../views/PrivateView.vue'
import PublicView from '../views/PublicView.vue'
import FrontView from '../views/FrontView.vue'
import FullMenu from '@/views/FullMenu.vue'
import OrderView from '@/views/OrderView.vue'
import OrderCompleteView from '@/views/OrderCompleteView.vue'


function getCookie(name) {
  const cookies = document.cookie.split(';');
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i].trim();    
    if (cookie.indexOf(name + '=') === 0) {
      return cookie.substring(name.length + 1);
    }
  }
  return null;
}

function checkAuth() {
  const token = getCookie('token');
  if (token) {
    return true;
  } else {
    return false;
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
      path: '/registrate',
      name: 'registrate',
      component: RegistrateView
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
    },
    {
      path: '/success', 
      name: 'success',
      component: OrderCompleteView
    }
  ],
  // scrollBehavior() {
  //   return { top: 0 };
  // }
})

router.beforeEach((to, from, next)=>{
  if ((to.path==='/private' && !checkAuth()) || (to.path==='/order' && !checkAuth())){
    next('/login')
  }else{
    next()
  }
})

export default router
