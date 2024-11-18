using System;
using Avalonia.Controls;
using Avalonia.Interactivity;
using FishingHellFrontend.HTTP;
using FishingHellFrontend.Views;
using Protobufs;

namespace FishingHellFrontend
{
    public partial class MainWindow : Window
    {
        private UserControl _fishingView;
        private UserControl _shopView;
        private UserControl _marketView;
        private UserControl _collectionView;
        private UserControl _loginView;
        public Calls _calls = new Calls();
        
        
        public Account? Account { get; set; }
        
        public MainWindow()
        {
            InitializeComponent();

            _fishingView = new FishingView(this);
            _shopView = new ShopView();
            _marketView = new MarketView();
            _collectionView = new CollectionView();
            _loginView = new LoginView(this);
            
            
            
            MainContent.Content = _loginView;
        }

        public void OnFishClicked(object sender, RoutedEventArgs e)
        {
            if (Account != null)
            {
                MainContent.Content = _fishingView;
                _fishingView.Focus(); // Ensure the view is ready to interact
            }
        }


        private void OnShopClicked(object sender, RoutedEventArgs e)
        {
            if (Account != null)
            {
                MainContent.Content = _shopView;
            }
        }

        private void OnMarketClicked(object sender, RoutedEventArgs e)
        {
            if (Account != null)
            {
                MainContent.Content = _marketView; 
            }
        }

        private void OnCollectionClicked(object sender, RoutedEventArgs e)
        {
            if (Account != null)
            {
                MainContent.Content = _collectionView; 
            }
        }
    }
}