<script setup>
    defineProps(['id', 'name', 'description', 'price'])
</script>
<template>
  <div class="card shadow rounded my-3">
            <div class="card-body">
                <div class="row">
                    <div class="col-md-2">
                        <img :src=this.methods.getPicture() class="product-image">
                    </div>
                    <div class="col-md-10 d-flex flex-column">
                        <div class="product-name">{{ name }}</div>
                        <div class="product-description">{{description}}</div>
                        <div class="flex-grow-1"></div>
                        <div class="d-flex justify-content-end">
                            <div class="product-price">Цена: <strong>{{ price / 100 }}</strong></div>
                            <button class="btn btn-warning" title="В корзину" onclick="alert('Товар добавлен в корзину')">
                                <i class="bi bi-cart"></i> <span class="d-none d-md-inline">В
                                    корзину</span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
</template>
<script>
export default {
    data() {
    return {
      methods: {
        getPicture: () => {
            return "assets/pics/product" + this.$props.id + ".jpg"
        }
      }
    }
  },
  computed: { 
  },
  methods: {
    
  },
  mounted() {
    console.log('Компонент примонтирован!');
  }
}
</script>
<style scoped>
</style>
