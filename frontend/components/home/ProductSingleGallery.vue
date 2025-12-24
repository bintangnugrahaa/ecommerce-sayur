<template>
    <div class="lg:w-1/2">
        <!-- Image slide -->
        <div class="product mb-6">
            <Swiper
                :modules="swiperModules"
                :navigation="true"
                :pagination="{ clickable: true }"
                :thumbs="{ swiper: thumbsSwiper }"
                @swiper="setMainSwiper"
                class="product-images-slider"
            >
                <SwiperSlide>
                    <div 
                        class="zoom" 
                        @mousemove="handleZoom"
                        :style="{ backgroundImage: `url(${defaultImage})` }"
                    >
                        <img :src="defaultImage" :alt="product?.product_name || 'Product Image'" class="w-full" />
                    </div>
                </SwiperSlide>
                <SwiperSlide v-for="(child, index) in product?.child" :key="index">
                    <div 
                        class="zoom" 
                        @mousemove="handleZoom"
                        :style="{ backgroundImage: `url(${child.image || defaultImage})` }"
                    >
                        <img :src="child.image || defaultImage" :alt="`${product?.product_name || 'Product'} - ${child.weight}${product?.unit}`" class="w-full" />
                    </div>
                </SwiperSlide>
            </Swiper>
        </div>
        
        <!-- Product thumbnails -->
        <div class="product-tools">
            <Swiper
                :modules="swiperModules"
                :slides-per-view="4"
                :space-between="10"
                :free-mode="true"
                :watch-slides-progress="true"
                @swiper="setThumbsSwiper"
                class="product-images-slider-thumbs"
            >
                <SwiperSlide class="cursor-pointer">
                    <img :src="defaultImage" :alt="product?.product_name || 'Product Image'" class="w-full" />
                </SwiperSlide>
                <SwiperSlide 
                    v-for="(child, index) in product?.child" 
                    :key="index"
                    class="cursor-pointer"
                >
                    <img :src="child.image || defaultImage" :alt="`${product?.product_name || 'Product'} - ${child.weight}${product?.unit}`" class="w-full" />
                </SwiperSlide>
            </Swiper>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Navigation, Pagination, FreeMode, Thumbs } from 'swiper/modules'
import 'swiper/css'
import 'swiper/css/navigation'
import 'swiper/css/pagination'
import 'swiper/css/free-mode'
import 'swiper/css/thumbs'

const props = defineProps({
    product: {
        type: Object,
        required: true,
        default: () => ({
            product_name: '',
            child: [],
            unit: ''
        })
    }
})

const swiperModules = [Navigation, Pagination, FreeMode, Thumbs]
const thumbsSwiper = ref(null)
const mainSwiper = ref(null)

const setThumbsSwiper = (swiper) => {
    thumbsSwiper.value = swiper
}

const setMainSwiper = (swiper) => {
    mainSwiper.value = swiper
}

const handleZoom = (event) => {
    const zoomer = event.currentTarget
    const offsetX = event.offsetX
    const offsetY = event.offsetY
    const x = (offsetX / zoomer.offsetWidth) * 100
    const y = (offsetY / zoomer.offsetHeight) * 100
    zoomer.style.backgroundPosition = `${x}% ${y}%`
}

const defaultImage = computed(() => {
    return props.product?.image || '/images/placeholder.jpg'
})
</script>

<style scoped>
.zoom {
    position: relative;
    overflow: hidden;
    background-position: 50% 50%;
    background-size: cover;
}

.zoom img:hover {
    opacity: 0;
}

.zoom img {
    transition: opacity 0.5s;
    display: block;
    width: 100%;
}

:deep(.swiper-button-next),
:deep(.swiper-button-prev) {
    color: #16a34a;
    background-color: white;
    padding: 20px;
    border-radius: 50%;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

:deep(.swiper-pagination-bullet-active) {
    background-color: #16a34a;
}

.product-images-slider-thumbs {
    padding: 10px 0;
}

.product-images-slider-thumbs :deep(.swiper-slide) {
    opacity: 0.5;
    transition: 0.3s;
}

.product-images-slider-thumbs :deep(.swiper-slide-thumb-active) {
    opacity: 1;
}
</style>