import { createRouter, createWebHistory } from 'vue-router'

import AppLayout from '@/layouts/AppLayout.vue'
import SearchView from '@/views/SearchView.vue'
import TagsOverviewView from '@/views/TagsOverviewView.vue'
import FilesView from '@/views/FilesView.vue'
import SettingsView from '@/views/SettingsView.vue'
import GoodbyeView from '@/views/GoodbyeView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

import { bootApp } from '@/router/middleware/bootApp'
import FullscreenLayout from '@/layouts/FullscreenLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: { name: 'app' }
    },
    {
      path: '/app',
      name: 'app',
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
          path: 'files',
          name: 'files',
          component: FilesView
        },
        {
          path: 'tags',
          name: 'tags',
          component: TagsOverviewView
        },
        {
          path: 'settings',
          name: 'settings',
          component: SettingsView
        },
        {
          path: 'goodbye',
          name: 'goodbye',
          component: GoodbyeView
        }
      ]
    },
    {
      path: '/say/',
      component: FullscreenLayout,
      children: [
        {
          path: 'goodbye',
          name: 'goodbye',
          component: GoodbyeView
        }
      ]
    },
    { path: '/:pathMatch(.*)*', name: '404', component: NotFoundView }
  ]
})

export default router
