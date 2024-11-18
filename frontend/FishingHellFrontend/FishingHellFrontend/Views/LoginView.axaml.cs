using Avalonia.Controls;
using Avalonia.Interactivity;
using System;
using FishingHellFrontend.HTTP;
using Protobufs;

namespace FishingHellFrontend.Views
{
    public partial class LoginView : UserControl
    {
        public Account? Account;
        private MainWindow _mainWindow;
        public LoginView(MainWindow mainWindow)
        {
            InitializeComponent();
            _mainWindow = mainWindow;
        }

        private void LoginButton_Click(object sender, RoutedEventArgs e)
        {
            var loginUsername = LoginUsernameTextBox.Text;
            var loginPassword = LoginPasswordBox.Text;
            if (string.IsNullOrWhiteSpace(loginUsername) || string.IsNullOrWhiteSpace(loginPassword))
            {
                LoginErrorTextBox.Text = "Please fill all fields";
                return;
            }
            Account = _mainWindow._calls.Login(loginUsername, loginPassword).Result;
            if (string.IsNullOrWhiteSpace(Account.Username) || string.IsNullOrWhiteSpace(Account.Password))
            {
                LoginErrorTextBox.Text = "Could not find username or password";
                return;
            }
            _mainWindow.Account = this.Account;
            _mainWindow.OnFishClicked(sender, e);
        }

        private void CreateAccountButton_Click(object sender, RoutedEventArgs e)
        {
            // Access the username and password for account creation
            var createUsername = CreateUsernameTextBox.Text;
            var createPassword = CreatePasswordBox.Text;

            if (string.IsNullOrWhiteSpace(createUsername) || string.IsNullOrWhiteSpace(createPassword))
            {
                CreateErrorTextBox.Text = "Please fill all the information!";
                return;
            }
            string result = _mainWindow._calls.Create(createUsername, createPassword).Result;
            if (!string.IsNullOrWhiteSpace(result))
            {
                CreateErrorTextBox.Text = "Username already used";
            }
            else
            {
                CreateErrorTextBox.Text = "Successfully created account";
                CreateUsernameTextBox.Text = "";
                CreatePasswordBox.Text = "";
            }
        }
    }
}