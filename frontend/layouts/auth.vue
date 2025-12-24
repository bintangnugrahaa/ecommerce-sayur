<script setup lang="ts">
import Footer from '~/components/common/Footer.vue';

const route = useRoute()
const isSignup =  computed(() => route.path === '/auth/signin')

</script>

<template>
    <main>
        <div class="border-b shadow-sm">
			<nav class="relative md:flex flex-wrap w-full items-center py-3 text-center">
				<div class="container sm:flex justify-between items-center">
					<a class="inline-block mr-4 text-lg whitespace-no-wrap" href="/">
						<img src="~/assets/images/logo/freshcart-logo.svg" alt="" class="inline-block align-text-top" />
					</a>
					<span v-if="!isSignup" class="inline-block">
						Already have an account?
						<a href="signin" class="text-green-600">Sign in</a>
					</span>
				</div>
			</nav>
		</div>


        <slot/>
    </main>
    <Footer/>
</template>