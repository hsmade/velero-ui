// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: { path: "/schedules" },
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: '/schedules',
        name: 'schedules',
        component: () => import(/* webpackChunkName: "home" */ '@/views/Schedules.vue'),
      },
      {
        path: '/schedule/:name',
        name: 'schedule',
        component: () => import(/* webpackChunkName: "home" */ '@/views/Schedule.vue'),
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
      {
        path: '/storagelocations',
        name: 'storagelocations',
        component: () => import(/* webpackChunkName: "home" */ '@/views/StorageLocations.vue'),
      },
      {
        path: '/storagelocation/:name',
        name: 'storagelocation',
        component: () => import(/* webpackChunkName: "home" */ '@/views/StorageLocation.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
