import Vue from "vue";
import VueRouter, {RouteConfig} from "vue-router";
import axios from 'axios';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    name: "home",
    component: () => import(/* webpackChunkName: "home" */ '../views/Home.vue'),
    alias: ["/put_name/*"],
    props: route => {
      const name = route.path.split("/");
      console.log(route.path, name);
      return { putName: name[name.length - 1] || false };
    }
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.afterEach(async () => {
  await axios.post('/api/visit');
})

export default router;
