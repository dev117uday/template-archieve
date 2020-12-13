
<template>
	<div class="hello">
		<b-container>
			<!-- svg -->
			<b-row>
				<b-col>
					<center>
						<img src="../assets/signup.svg" id="home-svg" />
					</center>
				</b-col>
			</b-row>
			<!-- end svg -->
			<!-- ########### -->
			<!-- sign title -->
			<b-row>
				<b-col id="signup-title">SignUp</b-col>
			</b-row>
			<!-- end title -->
			<!-- ########### -->
			<!-- form sign in -->
			<b-row>
				<b-col>
					<form action>
						<p id="form-headings">Name :</p>
						<center>
							<input v-model="userDetails[0]" class="input-signup" type="text" />
						</center>
						<p id="form-headings">Email :</p>
						<center>
							<input v-model="userDetails[1]" class="input-signup" type="email" />
						</center>
						<p id="form-headings">Password :</p>
						<center>
							<input v-model="userDetails[2]" class="input-signup" type="password" />
						</center>

						<p id="form-headings">Confirm Password :</p>
						<center>
							<input v-model="userDetails[3]" class="input-signup" type="password" />
						</center>
					</form>
				</b-col>
			</b-row>
			<!-- end form  -->
			<br />
			<br />
			<b-row>
				<!-- back button -->
				<b-col>
					<center>
						<form action="/user">
							<button
								type="submit"
								@click="buttonManager()"
								:style="
                  buttonStatus
                    ? 'background :#F50057 ; color : white'
                    : 'background : grey'
                "
								class="button"
							>Back</button>
						</form>
					</center>
				</b-col>

				<!-- end nack button-->
				<!-- ################################################ -->
				<!-- signup button -->
				<b-col>
					<center>
						<form @submit="signupForm($event)">
							<button
								type="submit"
								@click="buttonManager()"
								:style="
                  buttonStatus
                    ? 'background :#6C63FF ; color : white'
                    : 'background : grey'
                "
								class="button"
							>SignUp</button>
						</form>
					</center>
				</b-col>
				<!-- end buttion-->
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
			userDetails: ["", "", "", ""]
		};
	},

	methods: {
		buttonManager: function() {
			this.buttonStatus = !this.buttonStatus;
		},
		signupForm: function(event) {
			event.preventDefault();
			if (!this.userDetails[0]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Enter User Name");
			} else if (!this.userDetails[1]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Please Enter email ID");
			} else if (!this.userDetails[2]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Please enter password");
			} else if (!this.userDetails[3]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Enter confirm password");
			} else if (this.userDetails[2] != this.userDetails[3]) {
				this.buttonStatus = !this.buttonStatus;
				alert("Password doesnt match");
			} else if (this.userDetails[2].length < 8) {
				this.buttonStatus = !this.buttonStatus;
				alert("Password should be 8 character long");
			} else if (this.userDetails[0].split(" ").length != 2) {
				this.buttonStatus = !this.buttonStatus;
				alert("Enter full name like : 'Uday Ydav'");
			} else {
				let xm = this.userDetails[0].split(" ");
				let xf = xm[0].trim().toLowerCase();
				xf = xf.charAt(0).toUpperCase() + xf.slice(1);
				let xl = xm[1].trim().toLowerCase();
				xl = xl.charAt(0).toUpperCase() + xl.slice(1);
				xf = xf + " " + xl;

				auth()
					.createUserWithEmailAndPassword(
						this.userDetails[1],
						this.userDetails[2]
					)
					.then(() => {
						auth()
							.signInWithEmailAndPassword(
								this.userDetails[1],
								this.userDetails[2]
							)
							.then(response => {
								auth()
									.currentUser.updateProfile({
										displayName: xf
									})
									.then(() => {
										let item = {
											uid: response.user.uid,
											email: response.user.email,
											name: xf
										};
										store
											.setItem(
												"user",
												JSON.stringify(item)
											)
											.then(data => {
												console.log(data)

												this.$store.dispatch(
													"changeAuthStatus"
												);
												alert(
													"SignUp and LogIn sucsess"
												);
												window.location.href = "app";
											})
											.catch(err => {
												this.buttonStatus = !this
													.buttonStatus;
												alert(
													`Code : ${err.code} :: ${err.message}`
												);
											});
									})
									.catch(function(err) {
										this.buttonStatus = !this.buttonStatus;
										alert(
											`Code : ${err.code} :: ${err.message}`
										);
									});
							})
							.catch(err => {
								this.buttonStatus = !this.buttonStatus;
								alert(`Code : ${err.code} :: ${err.message}`);
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
	#signup-title {
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
	.input-signup {
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
	#form-headings {
		margin-left: 10vw;
		-webkit-touch-callout: none;
		-webkit-user-select: none;
		-khtml-user-select: none;
		-moz-user-select: none;
		-ms-user-select: none;
		user-select: none;
	}
</style>
