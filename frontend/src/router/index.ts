import { createRouter, createWebHistory } from 'vue-router'

import AppLayout from '@/layouts/AppLayout.vue'
import SearchView from '@/views/SearchView.vue'
import AboutView from '@/views/AboutView.vue'
import StatusView from '@/views/StatusView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

import { bootApp } from '@/router/middleware/bootApp'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { name: 'search' },
      beforeEnter: bootApp,
      component: AppLayout,
      children: [
        {
          path: 'search',
          name: 'search',
          component: SearchView
        },
        {
          path: 'about',
          name: 'about',
          component: AboutView
        },
        {
          path: 'status',
          name: 'status',
          component: StatusView
        }
      ]
    },
    { path: '/:pathMatch(.*)*', name: '404', component: NotFoundView }
  ]
})

export default router
