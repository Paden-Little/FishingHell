using System;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using FishingHellFrontend.Views;
using Google.Protobuf;

namespace FishingHellFrontend.HTTP;

public class Calls
{
    private HttpClient httpClient = new HttpClient();
    
    public async Task<Protobufs.Account> Login(string user, string password)
    {
        Protobufs.Account account = new();
        account.Username = user;
        account.Password = password;
        var response = httpClient.PostAsync("http://localhost:4321/login", new ByteArrayContent(account.ToByteArray()));
        account = Protobufs.Account.Parser.ParseFrom(await response.Result.Content.ReadAsByteArrayAsync());
        return account;
    }

    public async Task<String> Create(string username, string password)
    {
        Protobufs.Account account = new();
        account.Username = username;
        account.Password = password;
        var response = httpClient.PostAsync("http://localhost:4321/createAccount", new ByteArrayContent(account.ToByteArray()));
        return response.Result.StatusCode.ToString();
    }

    public async Task<Protobufs.Fish> Fish(string userID, string pool, int poleID, int baitID)
    {
        Protobufs.StartFish fish = new();
        fish.UserID = userID;
        fish.Pool = pool;
        fish.PoleID = poleID;
        fish.BaitID = baitID;
        
        var response = httpClient.PostAsync("http://localhost:4321/fish", new ByteArrayContent(fish.ToByteArray()));
        
        return Protobufs.Fish.Parser.ParseFrom(await response.Result.Content.ReadAsByteArrayAsync());
    }
}