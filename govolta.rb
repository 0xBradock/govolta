# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Govolta < Formula
  desc ""
  homepage "https://github.com/Kmelow/govolta"
  version "0.1.3"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/Kmelow/govolta/releases/download/v0.1.3/govolta_Darwin_x86_64.tar.gz"
      sha256 "f9fb67e692d09ac2505080f4ae8a1a14aabe0a8d3bdac6b3aefacd635a598c13"

      def install
        bin.install "govolta"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/Kmelow/govolta/releases/download/v0.1.3/govolta_Darwin_arm64.tar.gz"
      sha256 "945f64498e84e87a3db6ab03ffeeb5668af5f79258476eb7860cf0f3c9f52978"

      def install
        bin.install "govolta"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/Kmelow/govolta/releases/download/v0.1.3/govolta_Linux_arm64.tar.gz"
      sha256 "bccccfdeef2cacd68f128f3f7e0ac457ae01258a3ed46fd72e6ded482968fca0"

      def install
        bin.install "govolta"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/Kmelow/govolta/releases/download/v0.1.3/govolta_Linux_x86_64.tar.gz"
      sha256 "d4ba0b0667cbed95d665e2b1d953a82882e65818f3c57f3a5b9a8bfe6a9a7f65"

      def install
        bin.install "govolta"
      end
    end
  end
end