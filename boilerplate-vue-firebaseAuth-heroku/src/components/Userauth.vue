<template>
	<div class="hello">
		<b-container>
			<!-- Svg -->
			<b-row>
				<b-col>
					<center>
						<img src="../assets/home.svg" id="home-svg" />
					</center>
				</b-col>
			</b-row>
			<!-- End Svg -->
			<!-- ################################################ -->
			<!-- Title home -->
			<b-row>
				<b-col id="home-fill">App Name</b-col>
			</b-row>
			<!-- title End -->
			<br />
			<br />
			<b-row>
				<!-- login -->
				<b-col>
					<center>
						<form action="/login">
							<button
								@click="buttonManager()"
								type="submit"
								:style="
                  buttonStatus
                    ? 'background :#6C63FF ; color : white'
                    : 'background : grey'
                "
								id="button"
							>LogIn</button>
						</form>
					</center>
				</b-col>
				<!-- end login -->
				<!-- ################################################ -->
				<!-- Sign-Up -->
				<b-col>
					<center>
						<form action="/signup">
							<button
								type="submit"
								@click="buttonManager()"
								:style="
                  buttonStatus
                    ? 'background :#F50057 ; color : white'
                    : 'background : grey'
                "
								id="button"
							>SignUp</button>
						</form>
					</center>
				</b-col>
				<!-- End sign up -->
			</b-row>
			<br />
		</b-container>
	</div>
</template>

<script>
export default {
	data() {
		return {
			buttonStatus: true
		};
	},
	methods: {
		buttonManager: function() {
			this.buttonStatus = !this.buttonStatus;
		}
	}
};
</script>

<style scoped>
	html,
	body {
		margin: 0;
		overflow-x: hidden;
		overflow-y: hidden;
	}
	h1 {
		color: black;
	}
	#home-svg {
		margin-top: 10vh;
		width: 350px;
		height: 350px;
	}
	#home-fill {
		text-align: center;
		margin-top: -100px;
		font-size: 80px;
		font-family: sans-serif;
		font-weight: 400;
		-webkit-touch-callout: none;
		-webkit-user-select: none;
		-khtml-user-select: none;
		-moz-user-select: none;
		-ms-user-select: none;
		user-select: none;
	}
	#button {
		border-radius: 10px;
		margin-top: 2vh;
		font-family: sans-serif;
		font-weight: 500;
		font-size: 20px;
		height: 50px;
		outline: none;
		width: 140px;
	}
</style>
