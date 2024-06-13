<script setup>
import { useBucketStore } from '@/stores/BucketStore';
import ProductItem from '@/components/ProductItem.vue';
const bucketStore = useBucketStore()
</script>

<template >
 <section id="order" class="story">
			<div class="container">
				<h1 class="story__title">
					Основываясь на ваших предпочтениях:
				</h1>
				<p class="offer__intro">
					
				</p>
		</div>
		<div class="story-content container container_full-height" id ="story-content">
				<ProductItem v-for="product in bucketStore.order" :key="product.id_dish" :product="product" />
		</div>
		<RouterLink to="/private" class="offer__btn btn">
			В корзину
		</RouterLink>
</section>
</template>

<style scoped>

    .story-content {
	padding: 20px;
	background: #F1F1F5 no-repeat right / 55%;
	height: 600px;
	display: flex;
}

.story-content__item {
	font-size: 16px;
	width: 401px;
	position: relative;
	padding-left: 22px;
	line-height: 1.01em;
	margin-bottom: 74px;
	line-height:20px
}

.story-content__item:before {
	content:'';
	display: block;
	width: 7px;
	height: 40px;
	background: #9E9EB7;
	position: absolute;
	left: 0;
}
.story__title {
	font-size: 60px;
	letter-spacing: 0.22em;
	line-height: 1.35em;
	background-color: rgba(196, 210, 247, 0);
}
.story{
    padding-top: 15vh;
		height: 130vh;
	}
    .story-content__list{
	margin: 0 auto;
}

</style>