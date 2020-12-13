<template>
	<div class="hello">
		<b-container>
			<!-- svg -->
			<b-row>
				<b-col>
					<center>
						<img src="../assets/login.svg" id="home-svg" />
					</center>
				</b-col>
			</b-row>
			<!-- end svg -->
			<!-- ####################### -->
			<!-- end -->
			<b-row>
				<b-col id="login-title">Login</b-col>
			</b-row>
			<!-- end login -->
			<!-- ################ -->
			<!-- login form -->
			<b-row>
				<b-col>
					<form>
						<center>
							<p>Email :</p>
							<input v-model="userDetails[0]" class="input-login" type="email" />
							<p>Password :</p>
							<input v-model="userDetails[1]" class="input-login" type="password" />
						</center>
					</form>
				</b-col>
			</b-row>
			<!-- end login -->
			<br />
			<br />
			<b-row>
				<!-- login button -->
				<b-col>
					<center>
						<form action="/user">
							<button
								type="submit"
								@click="buttonManager()"
								:style="
                  buttonStatus
                    ? 'background :#6C63FF ; color : white'
                    : 'background : grey'
                "
								class="button"
							>Back</button>
						</form>
					</center>
				</b-col>
				<!-- end login button-->
				<!-- ################################################ -->
				<!-- back button-->
				<b-col>
					<center>
						<form @submit="loginForm($event)">
							<button
								type="submit"
								@click="buttonManager()"
								:style="
                  buttonStatus
                    ? 'background :#F50057 ; color : white'
                    : 'background : grey'
                "
								class="button"
							>LogIn</button>
						</form>
					</center>
				</b-col>
				<!-- End back button -->
			</b-row>
		</b-container>
		<br />
		<br />
		<br />
		<br />
		<br />
	</div>
</template>
<script>
require("../firebaseConfig");
import { auth } from "firebase";
import localforage from "localforage";
var store = localforage.createInstance({
	name: "users"
});
export default {
	data() {
		return {
			buttonStatus: true,
			userDetails: ["", ""]
		};
	},
	methods: {
		buttonManager: function() {
			this.buttonStatus = !this.buttonStatus;
		},
		loginForm: function(event) {
			event.preventDefault();
			if (!this.userDetails[0]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Please enter Email");
			} else if (!this.userDetails[1]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Please enter password");
			} else if (this.userDetails[1].length < 8) {
				this.buttonStatus = !this.buttonStatus;
				alert("Password you entered is less than 8 character");
			} else {
				auth()
					.signInWithEmailAndPassword(
						this.userDetails[0],
						this.userDetails[1]
					)
					.then(response => {
						let body = {
							uid: response.user.uid,
							email: response.user.email,
							name: response.user.displayName
						};
						store
							.setItem("user", JSON.stringify(body))
							.then(() => {
								this.$store.dispatch("changeAuthStatus");
								alert("LogIn sucsess");
								window.location.href = "app";
							})
							.catch(err => {
								console.log(err);
							});
					})
					.catch(err => {
						this.buttonStatus = !this.buttonStatus;
						alert(`Code : ${err.code} :: ${err.message}`);
					});
			}
		}
	}
};
</script>

<style scoped>
	#home-svg {
		height: 200px;
		width: 200px;
		margin-top: 5vh;
	}
	#login-title {
		text-align: center;
		font-family: sans-serif;
		font-size: 50px;
		font-weight: 400;
		-webkit-touch-callout: none;
		-webkit-user-select: none;
		-khtml-user-select: none;
		-moz-user-select: none;
		-ms-user-select: none;
		user-select: none;
	}
	.input-login {
		margin-top: -20px;
		font-family: "Roboto", sans-serif;
		outline: 0;
		background: #f2f2f2;
		width: 80%;
		border: 0;
		margin: -4px 0 10px;
		padding: 15px;
		box-sizing: border-box;
		font-size: 14px;
	}
	.button {
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
