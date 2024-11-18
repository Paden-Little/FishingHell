using System.Collections.ObjectModel;
using System.Windows.Input;
using FishingHellFrontend.Models;
using ReactiveUI;

namespace FishingHellFrontend.ViewModels

{
    public class FishingViewModel : ReactiveObject
    {
        // ObservableCollections for OwnedPoles and OwnedBaits
        private ObservableCollection<Pole> _ownedPoles = new ObservableCollection<Pole>();
        public ObservableCollection<Pole> OwnedPoles
        {
            get => _ownedPoles;
            set => this.RaiseAndSetIfChanged(ref _ownedPoles, value);
        }

        private ObservableCollection<Bait> _ownedBaits = new ObservableCollection<Bait>();
        public ObservableCollection<Bait> OwnedBaits
        {
            get => _ownedBaits;
            set => this.RaiseAndSetIfChanged(ref _ownedBaits, value);
        }

        // Commands
        public ICommand LogoutCommand { get; }
        public ICommand ShopCommand { get; }
        public ICommand MarketCommand { get; }
        public ICommand CollectionCommand { get; }
        public ICommand FishCommand { get; }

        public FishingViewModel()
        {
            // Initialize OwnedPoles and OwnedBaits collections with dummy data for now
            OwnedPoles.Add(new Pole { Name = "Basic Fishing Pole" });
            OwnedPoles.Add(new Pole { Name = "Advanced Rod" });
            OwnedPoles.Add(new Pole { Name = "Super Rod" });

            OwnedBaits.Add(new Bait { Name = "Worm" });
            OwnedBaits.Add(new Bait { Name = "Shrimp" });
            OwnedBaits.Add(new Bait { Name = "Insect" });

            // Initialize commands with stub methods
            LogoutCommand = ReactiveCommand.Create(Logout);
            ShopCommand = ReactiveCommand.Create(OpenShop);
            MarketCommand = ReactiveCommand.Create(OpenMarket);
            CollectionCommand = ReactiveCommand.Create(OpenCollection);
            FishCommand = ReactiveCommand.Create(Fish);
        }

        // Stub methods for each command
        private void Logout()
        {
            // Logic to handle user logout
        }

        private void OpenShop()
        {
            // Logic to open shop view
        }

        private void OpenMarket()
        {
            // Logic to open market view
        }

        private void OpenCollection()
        {
            // Logic to open collection view
        }

        private void Fish()
        {
            // Logic to handle fishing action
        }
    }
}
