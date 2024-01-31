using System;
using System.Text.Json;
using Api.Dtos;
using Api.Exceptions;
using Api.Predictors;
using Microsoft.Extensions.Logging;

namespace Api.Services.Predictors.Implementations;

public class PlayerPredictor : Predictor
{
    private readonly ILogger<Predictor> _logger;
    public PlayerPredictor(ILogger<Predictor> logger)
    {
        _logger = logger;
    }
    
    public override int Predict(JsonElement input)
    {
        PlayerInput playerInput;
        try
        { 
            playerInput = input.Deserialize<PlayerInput>();
        }
        catch (Exception e)
        {
            throw new InvalidPredictionInputException();
        }
        
        _logger.LogDebug("Player prediction: {} x {}", playerInput.FirstPlayerId, playerInput.SecondPlayerId);
        return playerInput.FirstPlayerId;
    }
}