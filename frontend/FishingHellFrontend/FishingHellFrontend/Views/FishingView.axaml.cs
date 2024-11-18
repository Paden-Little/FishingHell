using Avalonia.Controls;
using Avalonia.Controls.Primitives;
using Avalonia.Interactivity;
using System;
using System.Collections.ObjectModel;
using Avalonia;
using Avalonia.Media;
using FishingHellFrontend.HTTP;
using Protobufs;

namespace FishingHellFrontend.Views
{
    public partial class FishingView : UserControl
    {
        private MainWindow _mainWindow;
        public string poolID;
        public int baitID;
        public int poleID;
        public string[] Pools = ["Clear Lake", "Green River"];
        

        public FishingView(MainWindow mainWindow)
        {
            InitializeComponent();

            _mainWindow = mainWindow;
        }

        private void OnFishClicked(object sender, RoutedEventArgs e)
        {
            var fish = _mainWindow._calls.Fish(_mainWindow.Account.Id, poolID, poleID, baitID).Result;
            FishInfoTextBlock.Text = "Name: " + fish.Name
                + "\nTaxonomy: " + fish.Taxonomy
                + "\nWeight: " + fish.Weight;
        }
        
        private void OnPolesClicked(object sender, RoutedEventArgs e)
        {
            var flyout = new Flyout
            {
                Placement = PlacementMode.Bottom,
                Content = CreatePoleButtons()
            };

            flyout.ShowAt(PolesButton);
            PolesButton.Focus();
        }
        
        private StackPanel CreatePoleButtons()
        {
            var stackPanel = new StackPanel
            {
                Background = new SolidColorBrush(Colors.LightGray) // Highlight background for debugging
            };

            foreach (var pole in _mainWindow.Account.Poles)
            {
                var button = new Button
                {
                    Content = pole.Name,
                    Margin = new Thickness(5),
                    Padding = new Thickness(5),
                    Width = 100,
                    Height = 40
                };

                button.Click += (s, e) => EquipPole(pole);
                stackPanel.Children.Add(button);
            }

            return stackPanel;
        }

        private void EquipPole(Protobufs.Pole pole)
        {
            poleID = pole.Id;
            PolesButton.Content = "Pole: \n" + pole.Name;
        }

        private void OnBaitClicked(object sender, RoutedEventArgs e)
        {
            var flyout = new Flyout
            {
                Placement = PlacementMode.Bottom,
                Content = CreateBaitButtons()
            };

            flyout.ShowAt(BaitButton);
            BaitButton.Focus();
        }
        
        private StackPanel CreateBaitButtons()
        {
            var stackPanel = new StackPanel
            {
                Background = new SolidColorBrush(Colors.LightGray) // Highlight background for debugging
            };

            foreach (var bait in _mainWindow.Account.Baits)
            {
                var button = new Button
                {
                    Content = bait.Name + " [" + bait.Owned + "]",
                    Margin = new Thickness(5),
                    Padding = new Thickness(5),
                    Width = 100,
                    Height = 40
                };

                button.Click += (s, e) => EquipBait(bait);
                stackPanel.Children.Add(button);
            }

            return stackPanel;
        }
        
        private void EquipBait(Bait bait)
        {
            baitID = bait.Id;
            BaitButton.Content = "Bait: \n" + bait.Name + "[" + bait.Owned + "]";
        }

        private void OnPoolClicked(object sender, RoutedEventArgs e)
        {
            var flyout = new Flyout
            {
                Placement = PlacementMode.Bottom,
                Content = CreatePoolButtons()
            };

            flyout.ShowAt(PoolButton);
            PolesButton.Focus();
        }
        private StackPanel CreatePoolButtons()
        {
            var stackPanel = new StackPanel
            {
                Background = new SolidColorBrush(Colors.LightGray) // Highlight background for debugging
            };

            foreach (var pool in Pools)
            {
                var button = new Button
                {
                    Content = pool,
                    Margin = new Thickness(5),
                    Padding = new Thickness(5),
                    Width = 100,
                    Height = 40
                };

                button.Click += (s, e) => EquipPool(pool);
                stackPanel.Children.Add(button);
            }

            return stackPanel;
        }
        private void EquipPool(string pool)
        {
            poolID = pool;
            PoolButton.Content = "Pool: \n" + pool;
        }
    }
}
