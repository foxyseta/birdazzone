import { createRouter, createWebHistory } from 'vue-router';
import ChessView from '../views/ChessView.vue';
import FantaView from '../views/FantaView.vue';
import TvGamesView from '../views/TvGamesView.vue';
import TvGameView from '../views/TvGameView.vue';
import NotFoundView from '../views/NotFoundView.vue';

export const GAMES_PATH = '/tv-games';
export const CHESS_PATH = '/chess';
export const FANTA_PATH = '/fantacitorio';
export const GAME_ID_PATH = '/tv-games/:id';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Root',
      component: TvGamesView,
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: NotFoundView,
    },
    {
      path: GAMES_PATH,
      component: TvGamesView,
    },
    {
      path: CHESS_PATH,
      component: ChessView,
    },
    {
      path: FANTA_PATH,
      component: FantaView,
    },
    {
      path: GAME_ID_PATH,
      component: TvGameView,
      props: true,
    },
  ],
});

export default router;
