import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useBucketStore = defineStore('bucket', () => {

  let bucket = ref([
    {
        id_dish: 1,
        dish_name: 'Food for real 1',
        dish_image_path: '../src/assets/food.png',
        price: 420, 
        tags_id: 1
    },
    {
        id_dish: 2,
        dish_name: 'Food for real 2',
        dish_image_path: '../src/assets/food.png',
        price: 420, 
        tags_id: 1
    },
    {
        id_dish: 3,
        dish_name: 'Food for real 3',
        dish_image_path: '../src/assets/food.png',
        price: 420, 
        tags_id: 1
    },
    {
        id_dish: 4,
        dish_name: 'Food for real 4',
        dish_image_path: '../src/assets/food.png',
        price: 420, 
        tags_id: 1
    }, 
    {
        id_dish: 5,
        dish_name: 'water for real 1',
        dish_image_path: '../src/assets/water.png',
        price: 420, 
        tags_id: 2
    },
    {
        id_dish: 6,
        dish_name: 'water for real 2',
        dish_image_path: '../src/assets/water.png',
        price: 420, 
        tags_id: 2
    },
    {
        id_dish: 7,
        dish_name: 'water for real 3',
        dish_image_path: '../src/assets/water.png',
        price: 420, 
        tags_id: 2
    },
    {
        id_dish: 8,
        dish_name: 'water for real 4',
        dish_image_path: '../src/assets/water.png',
        price: 420, 
        tags_id: 2
    }, 
    {
        id_dish: 9,
        dish_name: 'fries for real 1',
        dish_image_path: '../src/assets/fries.jpg',
        price: 420, 
        tags_id: 3
    },
    {
        id_dish: 10,
        dish_name: 'fries for real 2',
        dish_image_path: '../src/assets/fries.jpg',
        price: 420, 
        tags_id: 3
    },
    {
        id_dish: 11,
        dish_name: 'fries for real 3',
        dish_image_path: '../src/assets/fries.jpg',
        price: 420, 
        tags_id: 3
    },
    {
        id_dish: 12, 
        dish_name: 'fries for real 4',
        dish_image_path: '../src/assets/fries.jpg',
        price: 420, 
        tags_id: 3
    }, 
   ])
   let order = ref([])
   const randomFood = localStorage.getItem('randomFood');
   if (randomFood) {
     order.value = JSON.parse(randomFood);
     console.log(bucket.value)
   }
  let buckLength = ref(bucket.value.length)
  let BuckExist = ref(bucket.value.length); 
  let totPrice = ref(0);
   bucket.value.forEach(element => {
       totPrice.value+=element.lastPrice
   });


   function addToBucket(id, title, lastPrice, oldPrice, image, count){
    let findIndex = bucket.value.findIndex(elem => elem.id === id)
    if (findIndex !== -1){
        
        // localStorage.bucket.value = JSON.stringify(bucket.value)
        bucket.value.splice(findIndex, 1)
        buckLength.value--
        localStorage.setItem('randomFood', JSON.stringify(bucket.value))
        console.log(bucket.value)
        // console.log(localStorage.getItem(JSON.stringify(('bucket.value'))))
        // console.log(bucket.value)
    }else{
        bucket.value.push({id: id, title: title, lastPrice: lastPrice, oldPrice: oldPrice, count: count,
            image: image })
        buckLength.value++
            localStorage.randomFood = JSON.stringify(bucket.value)
            console.log(bucket.value)
        //    console.log(localStorage.getItem(JSON.stringify('bucket.value')))
    }
    totPrice.value=0
    bucket.value.forEach(element => {
    totPrice.value+=element.lastPrice
    // BuckExist.value = BuckExist.value - 1
    });
    // if (findIndex !== -1){
    //     bucket.value = bucket.value.filter(elem => elem.id !== id)
    // }else{
    //     bucket.value.push({id: id, title: title, price: price,
    //         image: image, rating: rating, discount: discount, is_available: is_available})
    //        console.log(bucket.value)
    // }

    // bucket.value.push({id: id, title: title, price: price,
    //     image: image, rating: rating, discount: discount, is_available: is_available})
    //    console.log(bucket.value)
}


   function deleteItem(id){
       bucket.value = bucket.value.filter((item)=> +item.id != id)
       totPrice.value=0
       bucket.value.forEach(element => {
       totPrice.value+=element.lastPrice
       // BuckExist.value = BuckExist.value - 1
       });
       buckLength.value--
       localStorage.randomFood = JSON.stringify(bucket.value)
       console.log(id)
       console.log(buckLength.value)
       console.log('length ' + BuckExist.value)


       
   }

   function addRandomItem(tag){
    let neededProducts = bucket.value.filter(elem => elem.tags_id == tag)
    let item = neededProducts[Math.floor(Math.random()*neededProducts.length)];
    console.log('Added product with tag_id ' + tag + ' and id of product ' + item.id_dish)
    order.value.push(item)
    localStorage.randomFood = JSON.stringify(order.value)
    }


  return { bucket, buckLength, BuckExist, totPrice, order, deleteItem, addToBucket, addRandomItem }
})
