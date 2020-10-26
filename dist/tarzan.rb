# This file was generated by GoReleaser. DO NOT EDIT.
class Tarzan < Formula
  desc "Cross-platform shared clipboard"
  homepage "https://github.com/quackduck/tarzan"
  version "0.1.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/quackduck/tarzan/releases/download/v0.1.0/tarzan_0.1.0_Darwin_x86_64.tar.gz"
    sha256 "9af4c3bb956880d1eb412128eb04eeb652c8696b1a489769dc7f7e8d12bb27d7"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/quackduck/tarzan/releases/download/v0.1.0/tarzan_0.1.0_Linux_x86_64.tar.gz"
      sha256 "6683baad9670119a6c14be7f340f6602f2130d28f62a10903d83ad309d61c9b0"
    end
    if Hardware::CPU.arm?
      if Hardware::CPU.is_64_bit?
        url "https://github.com/quackduck/tarzan/releases/download/v0.1.0/tarzan_0.1.0_Linux_arm64.tar.gz"
        sha256 "162b6861137c7a06add534132b14d996eee9e9e9b6e46fd3db8246811dbc80ac"
      else
        url "https://github.com/quackduck/tarzan/releases/download/v0.1.0/tarzan_0.1.0_Linux_armv6.tar.gz"
        sha256 "8e2805d636ee15889d507d63ce0b3feec8cb31bb0f2e0cdd0ebc056822c57c51"
      end
    end
  end

  def install
    bin.install "tarzan"
  end
end