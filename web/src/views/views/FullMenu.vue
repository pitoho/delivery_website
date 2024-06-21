<script setup>
    import { useBucketStore } from '@/stores/BucketStore';
    import ProductItem from '../components/ProductItem.vue'

    const BucketStore = useBucketStore()
    let products = BucketStore.bucket
</script>

<template>
    <div class="product-list">
    <ProductItem v-for="product in products" :key="product.id_dish" :product="product" />
  </div>
</template>

<style scoped>
  body{
    overflow: scroll;
  }
    .product-list {
    overflow-y: scroll;
    padding-top: 10vh;

    overflow: scroll;
	max-height: 90vh;

  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
</style>