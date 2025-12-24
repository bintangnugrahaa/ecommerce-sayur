<template>
    <section class="container">
        <!-- row -->
        <div class="grid grid-cols-12 mb-8">
          <div class="col-span-12">
            <!-- card -->
            <div class="card bg-gray-100 border-0 rounded-2xl"
              style="background-image: url(/images/slider/slider-image-1.jpg); background-repeat: no-repeat; background-size: cover; background-position: right">
              <div class="card-body lg:p-12">
                <h1>Welcome back! FreshCart</h1>
                <p class="mb-6">FreshCart is simple & clean design for developer and designer.</p>
                <a href="#"
                  class="btn inline-flex items-center gap-x-2 bg-green-600 text-white border-green-600 disabled:opacity-50 disabled:pointer-events-none hover:text-white hover:bg-green-700 hover:border-green-700 active:bg-green-700 active:border-green-700 focus:outline-none focus:ring-4 focus:ring-green-100">
                  Create Product
                </a>
              </div>
            </div>
          </div>
        </div>
        <!-- table -->
    </section>
</template>

<script setup>
definePageMeta({
  middleware: ['admin', 'auth'],
  layout: 'dashboard',
})

useHead({
  script: [
  	{
		src: '/libs/simplebar/dist/simplebar.min.js',
		body: true
    },
    {
		src: '/js/theme.min.js',
		body: true
    },
    {
		src: '/libs/apexcharts/dist/apexcharts.min.js',
		body: true
    },
	{
		src: '/js/vendors/chart.js',
		body: true
    },
  ],
})
</script>