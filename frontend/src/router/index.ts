import { createRouter, createWebHistory } from 'vue-router'

import AppLayout from '@/layouts/AppLayout.vue'
import SearchView from '@/views/SearchView.vue'
import TagsOverviewView from '@/views/TagsOverviewView.vue'
import FilesView from '@/views/FilesView.vue'
import SettingsView from '@/views/SettingsView.vue'
import GoodbyeView from '@/views/GoodbyeView.vue'
import OtpAuthView from '@/views/OtpAuthView.vue'
import NotFoundView from '@/views/NotFoundView.vue'

import { bootApp } from '@/router/middleware/bootApp'
import FullscreenLayout from '@/layouts/FullscreenLayout.vue'

declare module 'vue-router' {
  interface RouteMeta {
    title?: string
  }
}

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
          component: SearchView,
          meta: { title: 'menu.search.label' }
        },
        {
          path: 'files',
          name: 'files',
          component: FilesView,
          meta: { title: 'menu.files.label' }
        },
        {
          path: 'tags',
          name: 'tags',
          component: TagsOverviewView,
          meta: { title: 'menu.tags.label' }
        },
        {
          path: 'settings',
          name: 'settings',
          component: SettingsView,
          meta: { title: 'menu.settings.label' }
        },
        {
          path: 'goodbye',
          name: 'goodbye',
          component: GoodbyeView,
          meta: { title: 'menu.quit.label' }
        }
      ]
    },
    {
      path: '/auth',
      component: FullscreenLayout,
      props: { showLogo: false },
      children: [
        {
          path: 'otp',
          name: 'otp',
          component: OtpAuthView
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
