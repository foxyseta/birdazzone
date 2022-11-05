import { createRouter, createWebHistory } from 'vue-router'
import EreditaView from '../views/EreditaView.vue'
import CatenaView from '../views/CatenaView.vue'
import ChessView from '../views/ChessView.vue'
import TvGamesView from '../views/TvGamesView.vue'
import NotFoundView from '../views/NotFoundView.vue'

export const EREDITA_PATH = "/games/leredita"
export const GAMES_PATH = "/games"
export const CATENA_PATH = "/games/catena"
export const CHESS_PATH = "/chess"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Root',
      component: EreditaView
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: NotFoundView
    },
    {
      path: GAMES_PATH,
      component: TvGamesView
    },
    {
      path: EREDITA_PATH,
      component: EreditaView
    },
    {
      path: CATENA_PATH,
      component: CatenaView
    },
    {
      path: CHESS_PATH,
      component: ChessView
    },
  ]
})

export default router
