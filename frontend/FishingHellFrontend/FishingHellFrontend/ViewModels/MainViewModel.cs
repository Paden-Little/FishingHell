using CommunityToolkit.Mvvm.ComponentModel;

namespace FishingHellFrontend.ViewModels;

public partial class MainViewModel : ViewModelBase
{
    [ObservableProperty] private string _greeting = "Welcome to GAY Avalonia!";
}