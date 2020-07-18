/* eslint-disable */
import Vue from 'vue';
import Router from 'vue-router';
import Ping from "../components/Ping";
import Rooms from "@/components/Rooms";
Vue.use(Router);

export default new Router({
  routes: [

    {
      path: '/ping',
      name: 'Ping',
      component: Ping,
},
    {
      path: '/',
      name: 'Rooms',
      component: Rooms,
    },
  ], mode: 'history',
});
