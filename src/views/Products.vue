<script setup>
import Product from './Product.vue'
</script>

<template>
  <html>

  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Товары</title>
  </head>

  <body>
    <main class="container">
      <ul v-for="(product, index) in products" :key="index">
        <Product v-bind:name=product.name v-bind:price=product.price v-bind:description=product.description
          v-bind:id=product.id> </Product>
      </ul>
      <ul class="list-group list-group-horizontal-sm">
        <li class="list-group-item" v-for="page in pages"> <button type="button" class="btn btn-link" @click="methods.selectPage(page - 1)">{{ page }}</button></li>
      </ul>

    </main>
    <footer>
    </footer>
  </body>

  </html>
</template>

<script>
export default {
  data() {
    return {
      currentRating: 0,
      products: null,
      pages: 0,
      methods: {
        selectPage: (num) => {
          fetch("http://localhost:1323/products/get?page=" + num)
            .then(r => r.json())
            .then(json => {
              this.products = json
            });
        }
      }
    }
  },
  computed: {
  },
  methods: {
  },
  mounted() {
    fetch("http://localhost:1323/products/get?page=0")
      .then(r => r.json())
      .then(json => {
        this.products = json
      });

    fetch("http://localhost:1323/products/pages")
      .then(r => r.json())
      .then(json => {
        this.pages = json
      });

    console.log('Компонент примонтирован!');
  }
}
</script>
