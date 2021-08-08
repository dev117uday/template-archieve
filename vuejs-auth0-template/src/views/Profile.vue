<template>
  <div class="container mt-5">
    <div class="row g-0">
      <div class="col-md-4">
        <img :src="$auth.user.picture" id="profileimg" />
      </div>
      <div class="col-md-8">
        <div class="card-body">
          <h3 class="card-title">{{ $auth.user.name }}</h3>
          <p class="card-text">{{ $auth.user.email }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
#profileimg {
  max-width: 300px;
  max-height: 300px;
  border-radius: 10%;
}
</style>
