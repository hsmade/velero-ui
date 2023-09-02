// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '',
        name: 'home',
        component: () => import(/* webpackChunkName: "home" */ '@/views/Home.vue'),
      },
      {
        path: '/backups',
        name: 'backups',
        component: () => import(/* webpackChunkName: "home" */ '@/views/Backups.vue'),
      },
      {
        path: '/backup/:name',
        name: 'backup',
        component: () => import(/* webpackChunkName: "home" */ '@/views/Backup.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
